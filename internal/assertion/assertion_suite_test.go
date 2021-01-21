package assertion_test

import (
	. "github.com/ptcar2009/ginkgo"
	. "github.com/ptcar2009/gomega"

	"testing"
)

func TestAssertion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Assertion Suite")
}
