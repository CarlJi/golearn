package e2e

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/qbox/monos/test/framework"
)

var _ = Describe("your service", func() {
	f := framework.NewDefaultFramework("servicename")

	It("should do", func() {
		Expect(f.Namespace.Name).To(ContainSubstring("e2e"))
	})

	It("should not do", func() {
		Expect(f.Namespace.Name).To(ContainSubstring("e2e"))
	})
})
