package content

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestContent(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Content Suite")
}
