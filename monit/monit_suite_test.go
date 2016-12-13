package monit

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMonit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Monit Suite")
}

func readFile(path string) []byte {
	contents, err := ioutil.ReadFile(path)
	Expect(err).NotTo(HaveOccurred())
	return contents
}
