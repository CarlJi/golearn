package framework

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang/glog"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
)

const (
	// How often to Poll pods, nodes and claims.
	Poll = 2 * time.Second

	// service accounts are provisioned after namespace creation
	// a service account is required to support pod creation in a namespace as part of admission control
	ServiceAccountProvisionTimeout = 2 * time.Minute

	DefaultNamespaceDeletionTimeout = 5 * time.Minute
)

func CreateTestingNS(baseName string, c clientset.Interface, lables map[string]string) (*v1.Namespace, error) {
	if lables == nil {
		lables = make(map[string]string)
	}

	lables["e2e-run"] = string(RunId)

	namespaceObj := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: fmt.Sprintf("e2e-test-%v-", baseName),
			Namespace:    "",
			Labels:       lables,
		},
		Status: v1.NamespaceStatus{},
	}

	var got *v1.Namespace

	if err := wait.PollImmediate(Poll, 30*time.Second, func() (bool, error) {
		var err error
		got, err = c.CoreV1().Namespaces().Create(namespaceObj)
		if err != nil {
			glog.Warningf("unexpected error when creating namespace: %v\n", err)
			return false, nil
		}
		return true, nil
	}); err != nil {
		return nil, err
	}

	if err := WaitForDefaultAccountInNamespace(c, got.Name); err != nil {
		return nil, err
	}

	return got, nil
}

func deleteNS(c clientset.Interface, clientPool dynamic.ClientPool, namespace string, timeout time.Duration) error {
	startTime := time.Now()
	if err := c.CoreV1().Namespaces().Delete(namespace, nil); err != nil {
		return err
	}

	err := wait.PollImmediate(2*time.Second, timeout, func() (bool, error) {
		if _, err := c.CoreV1().Namespaces().Get(namespace, metav1.GetOptions{}); err != nil {
			if apierrs.IsNotFound(err) {
				return true, nil
			}
			Logf("Error while waiting for namespace to be terminated: %v", err)
			return false, nil
		}
		return false, nil
	})

	// verify there is no more remaining content in the namespace
	remainingContent, cerr := hasRemainingContent(c, clientPool, namespace)
	if cerr != nil {
		return cerr
	}

	// if content remains, let's dump information about the namespace, and system for flake debugging.
	remainingPods := 0
	missingTimestamp := 0
	if remainingContent {
		// log information about namespace, and set of namespaces in api server to help flake detection
		logNamespace(c, namespace)
		logNamespaces(c, namespace)

		// if we can, check if there were pods remaining with no timestamp.
		remainingPods, missingTimestamp, _ = countRemainingPods(c, namespace)
	}

	// a timeout waiting for namespace deletion happened!
	if err != nil {
		// some content remains in the namespace
		if remainingContent {
			// pods remain
			if remainingPods > 0 {
				if missingTimestamp != 0 {
					// pods remained, but were not undergoing deletion (namespace controller is probably culprit)
					return fmt.Errorf("namespace %v was not deleted with limit: %v, pods remaining: %v, pods missing deletion timestamp: %v", namespace, err, remainingPods, missingTimestamp)
				}
				// but they were all undergoing deletion (kubelet is probably culprit, check NodeLost)
				return fmt.Errorf("namespace %v was not deleted with limit: %v, pods remaining: %v", namespace, err, remainingPods)
			}
			// other content remains (namespace controller is probably screwed up)
			return fmt.Errorf("namespace %v was not deleted with limit: %v, namespaced content other than pods remain", namespace, err)
		}
		// no remaining content, but namespace was not deleted (namespace controller is probably wedged)
		return fmt.Errorf("namespace %v was not deleted with limit: %v, namespace is empty but is not yet removed", namespace, err)
	}
	Logf("namespace %v deletion completed in %s", namespace, time.Now().Sub(startTime))
	return nil

}

func hasRemainingContent(c clientset.Interface, clientPool dynamic.ClientPool, namespace string) (bool, error) {
	if clientPool == nil {
		return false, nil
	}

	resources, err := c.Discovery().ServerPreferredNamespacedResources()
	if err != nil {
		return false, err
	}

	groupVersionResources, err := discovery.GroupVersionResources(resources)
	if err != nil {
		return false, nil
	}

	ignoreResources := sets.NewString("bindings")

	contentRemaining := false

	for gvr := range groupVersionResources {
		dynamicClient, err := clientPool.ClientForGroupVersionResource(gvr)
		if err != nil {
			// not all resource types support list, so some errors here are normal depending on the resource type.
			Logf("namespace: %s, unable to get client - gvr: %v, error: %v", namespace, gvr, err)
			continue
		}

		apiResource := metav1.APIResource{Name: gvr.Resource, Namespaced: true}
		// TODO: temporary hack for https://github.com/kubernetes/kubernetes/issues/31798
		if ignoreResources.Has(apiResource.Name) {
			Logf("namespace: %s, resource: %s, ignored listing per whitelist", namespace, apiResource.Name)
			continue
		}

		obj, err := dynamicClient.Resource(&apiResource, namespace).List(metav1.ListOptions{})
		if err != nil {
			// not all resources support list, so we ignore those
			if apierrs.IsMethodNotSupported(err) || apierrs.IsNotFound(err) || apierrs.IsForbidden(err) {
				continue
			}
			return false, err
		}

		unstructuredList, ok := obj.(*unstructured.UnstructuredList)
		if !ok {
			return false, fmt.Errorf("namespace: %s, resource: %s, expected *unstructured.UnstructuredList, got %#v", namespace, apiResource.Name, obj)
		}
		if len(unstructuredList.Items) > 0 {
			Logf("namespace: %s, resource: %s, items remaining: %v", namespace, apiResource.Name, len(unstructuredList.Items))
			contentRemaining = true
		}
	}
	return contentRemaining, nil
}

// countRemainingPods queries the server to count number of remaining pods, and number of pods that had a missing deletion timestamp.
func countRemainingPods(c clientset.Interface, namespace string) (int, int, error) {
	// check for remaining pods
	pods, err := c.Core().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return 0, 0, err
	}

	// nothing remains!
	if len(pods.Items) == 0 {
		return 0, 0, nil
	}

	// stuff remains, log about it
	logPodStates(pods.Items)

	// check if there were any pods with missing deletion timestamp
	numPods := len(pods.Items)
	missingTimestamp := 0
	for _, pod := range pods.Items {
		if pod.DeletionTimestamp == nil {
			missingTimestamp++
		}
	}
	return numPods, missingTimestamp, nil
}

// logPodStates logs basic info of provided pods for debugging.
func logPodStates(pods []v1.Pod) {
	// Find maximum widths for pod, node, and phase strings for column printing.
	maxPodW, maxNodeW, maxPhaseW, maxGraceW := len("POD"), len("NODE"), len("PHASE"), len("GRACE")
	for i := range pods {
		pod := &pods[i]
		if len(pod.ObjectMeta.Name) > maxPodW {
			maxPodW = len(pod.ObjectMeta.Name)
		}
		if len(pod.Spec.NodeName) > maxNodeW {
			maxNodeW = len(pod.Spec.NodeName)
		}
		if len(pod.Status.Phase) > maxPhaseW {
			maxPhaseW = len(pod.Status.Phase)
		}
	}
	// Increase widths by one to separate by a single space.
	maxPodW++
	maxNodeW++
	maxPhaseW++
	maxGraceW++

	// Log pod info. * does space padding, - makes them left-aligned.
	Logf("%-[1]*[2]s %-[3]*[4]s %-[5]*[6]s %-[7]*[8]s %[9]s",
		maxPodW, "POD", maxNodeW, "NODE", maxPhaseW, "PHASE", maxGraceW, "GRACE", "CONDITIONS")
	for _, pod := range pods {
		grace := ""
		if pod.DeletionGracePeriodSeconds != nil {
			grace = fmt.Sprintf("%ds", *pod.DeletionGracePeriodSeconds)
		}
		Logf("%-[1]*[2]s %-[3]*[4]s %-[5]*[6]s %-[7]*[8]s %[9]s",
			maxPodW, pod.ObjectMeta.Name, maxNodeW, pod.Spec.NodeName, maxPhaseW, pod.Status.Phase, maxGraceW, grace, pod.Status.Conditions)
	}
	Logf("") // Final empty line helps for readability.
}

// logNamespaces logs the number of namespaces by phase
// namespace is the namespace the test was operating against that failed to delete so it can be grepped in logs
func logNamespaces(c clientset.Interface, namespace string) {
	namespaceList, err := c.Core().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		Logf("namespace: %v, unable to list namespaces: %v", namespace, err)
		return
	}

	numActive := 0
	numTerminating := 0
	for _, namespace := range namespaceList.Items {
		if namespace.Status.Phase == v1.NamespaceActive {
			numActive++
		} else {
			numTerminating++
		}
	}
	Logf("namespace: %v, total namespaces: %v, active: %v, terminating: %v", namespace, len(namespaceList.Items), numActive, numTerminating)
}

// logNamespace logs detail about a namespace
func logNamespace(c clientset.Interface, namespace string) {
	ns, err := c.Core().Namespaces().Get(namespace, metav1.GetOptions{})
	if err != nil {
		if apierrs.IsNotFound(err) {
			Logf("namespace: %v no longer exists", namespace)
			return
		}
		Logf("namespace: %v, unable to get namespace due to error: %v", namespace, err)
		return
	}
	Logf("namespace: %v, DeletionTimetamp: %v, Finalizers: %v, Phase: %v", ns.Name, ns.DeletionTimestamp, ns.Spec.Finalizers, ns.Status.Phase)
}

func WaitForDefaultAccountInNamespace(c clientset.Interface, namespace string) error {
	return WaitForAccountInNamespace(c, namespace, "default", ServiceAccountProvisionTimeout)
}

func WaitForAccountInNamespace(c clientset.Interface, namespace, accountName string, timeout time.Duration) error {
	w, err := c.CoreV1().ServiceAccounts(namespace).Watch(metav1.SingleObject(metav1.ObjectMeta{Name: accountName}))
	if err != nil {
		return err
	}

	_, err = watch.Until(timeout, w, ServiceAccountHasSecrets)
	if err != nil {
		return err
	}

	return nil
}

// ServiceAccountHasSecrets returns true if the service account has at least one secret,
// false if it does not, or an error.
func ServiceAccountHasSecrets(event watch.Event) (bool, error) {
	switch event.Type {
	case watch.Deleted:
		return false, errors.New("serviceaccounts not found")
	}
	switch t := event.Object.(type) {
	case *v1.ServiceAccount:
		return len(t.Secrets) > 0, nil
	}
	return false, nil
}
