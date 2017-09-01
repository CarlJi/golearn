package framework

import (
	"fmt"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/client-go/dynamic"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type Framework struct {
	BaseName string

	// ClientSet uses internal objects, you should use ClientSet where possible.
	ClientSet  clientset.Interface
	ClientPool dynamic.ClientPool

	// configuration for framework's client
	Options               FrameworkOptions
	SkipNamespaceCreation bool // Whether to skip creating a namespace

	Namespace          *v1.Namespace   // Every test has at least one namespace unless creation is skipped
	namespacesToDelete []*v1.Namespace // Some tests have more than one.
}

type FrameworkOptions struct {
}

func NewDefaultFramework(baseName string) *Framework {
	return NewFramework(baseName, FrameworkOptions{}, nil)
}

func NewFramework(baseName string, options FrameworkOptions, client clientset.Interface) *Framework {
	f := &Framework{
		BaseName:  baseName,
		ClientSet: client,
		Options:   options,
	}

	BeforeEach(f.BeforeEach)
	AfterEach(f.AfterEach)

	return f
}

// BeforeEach gets a clientset and makes a namespace
func (f *Framework) BeforeEach() {
	if f.ClientSet == nil {
		By("Creating a kubernetes client")
		config, err := LoadConfig()
		Expect(err).NotTo(HaveOccurred())
		f.ClientSet, err = clientset.NewForConfig(config)
		Expect(err).NotTo(HaveOccurred())
	}

	if !f.SkipNamespaceCreation {
		ns, err := f.CreateNamespace(f.BaseName, map[string]string{
			"e2e-framework": f.BaseName,
		})
		Expect(err).NotTo(HaveOccurred(), "failed to create namespace")

		By(fmt.Sprintf("Create namespace %s successfully", ns.Name))

		f.Namespace = ns
	}
}

func (f *Framework) AfterEach() {
	defer func() {
		nsDelettionErrors := map[string]error{}
		// Whether to delete namespace is determined by 3 factors: delete-namespace flag, delete-namespace-on-failure flag and the test result
		// if delete-namespace set to false, namespace will always be preserved.
		// if delete-namespace is true and delete-namespace-on-failure is false, namespace will be preserved if test failed.
		if TestContext.DeleteNamespace && (TestContext.DeleteNamespaceOnFailure || !CurrentGinkgoTestDescription().Failed) {
			for _, ns := range f.namespacesToDelete {
				By(fmt.Sprintf("Destroying namespace %q for this suite", ns.Name))
				if err := deleteNS(f.ClientSet, f.ClientPool, ns.Name, DefaultNamespaceDeletionTimeout); err != nil {
					if !errors.IsNotFound(err) {
						nsDelettionErrors[ns.Name] = err
					} else {
						Logf("Namespace %v was already deleted", ns.Name)
					}
				}
			}
		} else {
			if !TestContext.DeleteNamespace {
				Logf("Found DeleteNamespace=false, skipping namespace deletion!")
			} else {
				Logf("Found DeleteNamespaceOnFailure=false and current test failed, skipping namespace deletion!")
			}
		}

		f.Namespace = nil
		f.ClientSet = nil
		f.namespacesToDelete = nil

		if len(nsDelettionErrors) > 0 {
			messages := []string{}
			for namesapceKey, namesapceErr := range nsDelettionErrors {
				messages = append(messages, fmt.Sprintf("Couldn't delete ns: %q: %s (%#v)", namesapceKey, namesapceErr, namesapceErr))
			}
			Fail(strings.Join(messages, ","))
		}
	}()
}

func (f *Framework) CreateNamespace(baseName string, labels map[string]string) (*v1.Namespace, error) {
	ns, err := CreateTestingNS(f.BaseName, f.ClientSet, labels)
	// check ns instead of err ro see if its nil as we may
	// fail to create serviceaccount in it.
	// In this case. we should not forget to delete the namespace
	if ns != nil {
		f.namespacesToDelete = append(f.namespacesToDelete, ns)
	}
	return ns, err
}

// unique identifier of the e2e run
var RunId = uuid.NewUUID()

func LoadConfig() (*restclient.Config, error) {
	c, err := RestclientConfig()
	if err != nil {
		if TestContext.KubeConfig == "" {
			return restclient.InClusterConfig()
		} else {
			return nil, err
		}
	}

	return clientcmd.NewDefaultClientConfig(*c, &clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: TestContext.Host}}).ClientConfig()
}

func RestclientConfig() (*clientcmdapi.Config, error) {
	Logf(">>> kubeConfig: %s", TestContext.KubeConfig)
	if TestContext.KubeConfig == "" {
		return nil, fmt.Errorf("KubeConfig must be specified to load client config")
	}
	c, err := clientcmd.LoadFromFile(TestContext.KubeConfig)
	if err != nil {
		return nil, fmt.Errorf("error loading KubeConfig: %v", err.Error())
	}
	return c, nil
}

func nowStamp() string {
	return time.Now().Format(time.StampMilli)
}

func log(level string, format string, args ...interface{}) {
	fmt.Fprintf(GinkgoWriter, nowStamp()+": "+level+": "+format+"\n", args...)
}

func Logf(format string, args ...interface{}) {
	log("INFO", format, args...)
}
