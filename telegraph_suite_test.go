package telegraph_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTelegraph(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Telegraph Suite")
}
