package emulator_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEmulator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Emulator Suite")
}
