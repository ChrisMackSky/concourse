package compression_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCompression(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Compression Suite")
}
