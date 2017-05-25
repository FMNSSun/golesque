package golesque_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGolesque(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Golesque Suite")
}
