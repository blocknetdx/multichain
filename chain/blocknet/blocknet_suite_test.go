package blocknet_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBlocknet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Blocknet Suite")
}
