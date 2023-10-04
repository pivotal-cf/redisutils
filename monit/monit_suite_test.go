package monit

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMonit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "monit suite")
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

func joinCommand(command string, args []string) string {
	joined := []string{command}
	return strings.Join(append(joined, args...), " ")
}

func combinedOutputReturns(returns []byteSliceAndError) *combinedOutput {
	return &combinedOutput{returnIndex: -1, returns: returns}
}

type byteSliceAndError struct {
	byteSlice []byte
	err       error
}

type combinedOutput struct {
	returns     []byteSliceAndError
	returnIndex int
}

func (c *combinedOutput) sequentially() ([]byte, error) {
	c.returnIndex++
	return c.returns[c.returnIndex].byteSlice, c.returns[c.returnIndex].err
}
