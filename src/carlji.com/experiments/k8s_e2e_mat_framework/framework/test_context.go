package framework

import (
	"flag"
	"fmt"

	"k8s.io/client-go/tools/clientcmd"
)

const defaultHost = "http://127.0.0.1:8080"

type TestContextType struct {
	KubeConfig string
	ReportDir  string
	Host       string

	DeleteNamespace          bool
	DeleteNamespaceOnFailure bool
}

var TestContext TestContextType

func RegisterFlags() {
	flag.StringVar(&TestContext.KubeConfig, clientcmd.RecommendedConfigPathFlag, clientcmd.RecommendedHomeFile, "Path to kubeconfig containing embedded authinfo.")
	flag.StringVar(&TestContext.ReportDir, "report-dir", "", "Path to the directory where the JUnit XML reports should be saved. Default is empty, which doesn't generate these reports.")
	flag.StringVar(&TestContext.Host, "host", "", fmt.Sprintf("The host, or apiserver, to connect to. Will default to %s if this argument and --kubeconfig are not set", defaultHost))
	flag.BoolVar(&TestContext.DeleteNamespace, "delete-namespace", true, "If true tests will delete namespace after completion. It is only designed to make debugging easier, DO NOT turn it off by default.")
	flag.BoolVar(&TestContext.DeleteNamespaceOnFailure, "delete-namespace-on-failure", false, "If true, framework will delete test namespace on failure. Used only during test debugging.")
}
