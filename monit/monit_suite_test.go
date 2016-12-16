package monit

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

func getExampleMonitSummary() []byte {
	path := filepath.FromSlash("assets/example_monit_summary.txt")
	return readFile(path)
}

func getExampleMonitSummaryOneStopped() []byte {
	path := filepath.FromSlash("assets/example_monit_summary_one_stopped.txt")
	return readFile(path)
}

func getExampleMonitSummaryAllStatuses() []byte {
	path := filepath.FromSlash("assets/example_monit_summary_all_statuses.txt")
	return readFile(path)
}
