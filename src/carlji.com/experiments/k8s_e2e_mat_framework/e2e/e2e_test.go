package e2e

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang/glog"
	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/gomega"
	"github.com/qbox/monos/test/framework"
)

func init() {
	framework.RegisterFlags()
}

func TestE2E(t *testing.T) {
	RunE2ETests(t)
}

func RunE2ETests(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)

	var r []ginkgo.Reporter
	var ReportDir = "reports"

	if framework.TestContext.ReportDir != "" {
		ReportDir = framework.TestContext.ReportDir
	}

	if err := os.Mkdir(ReportDir, os.ModePerm); err != nil && !os.IsExist(err) {
		glog.Fatalf("Failed creating report directory %s ", ReportDir)
	}

	r = append(r, reporters.NewJUnitReporter(filepath.Join(ReportDir, fmt.Sprintf("service_%02d.xml", config.GinkgoConfig.ParallelNode))))

	framework.Logf("Starting e2e run %q on ginkgo node %d \n", framework.RunId, config.GinkgoConfig.ParallelNode)
	ginkgo.RunSpecsWithDefaultAndCustomReporters(t, "e2e test suite", r)
}
