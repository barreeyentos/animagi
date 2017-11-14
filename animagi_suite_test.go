package animagi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAnimagi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Animagi Suite")
}
