package gstruct_test

import (
	"testing"

	. "github.com/ptcar2009/ginkgo"
	. "github.com/ptcar2009/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gstruct Suite")
}
