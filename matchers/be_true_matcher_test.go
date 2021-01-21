package matchers_test

import (
	. "github.com/ptcar2009/ginkgo"
	. "github.com/ptcar2009/gomega"
	. "github.com/ptcar2009/gomega/matchers"
)

var _ = Describe("BeTrue", func() {
	It("should handle true and false correctly", func() {
		Expect(true).Should(BeTrue())
		Expect(false).ShouldNot(BeTrue())
	})

	It("should only support booleans", func() {
		success, err := (&BeTrueMatcher{}).Match("foo")
		Expect(success).Should(BeFalse())
		Expect(err).Should(HaveOccurred())
	})
})
