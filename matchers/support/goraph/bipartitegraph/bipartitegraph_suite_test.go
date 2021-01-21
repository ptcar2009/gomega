package bipartitegraph_test

import (
	. "github.com/ptcar2009/ginkgo"
	. "github.com/ptcar2009/gomega"

	"testing"
)

func TestBipartitegraph(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bipartitegraph Suite")
}
