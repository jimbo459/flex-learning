package interpreter

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInterpreter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interpreter Suite")
}
