package format_test

import (
	. "github.com/ptcar2009/ginkgo"
	. "github.com/ptcar2009/gomega"

	"testing"
)

func TestFormat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Format Suite")
}
