package monit_integration_test

import (
	"fmt"
	"os/exec"
	"regexp"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMonitIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Monit Integration Suite")
}

var _ = BeforeSuite(func() {
	startMonit()
	waitForMonitJobsToBeRunning()
})

func waitForMonitJobsToBeRunning() {
	waitFor(fooIsRunning)
	waitFor(barIsRunning)
	waitFor(bazIsRunning)
}

func waitFor(isDone func() bool) {
	interval := time.Millisecond * 200
	timeout := time.Second * 15

	for elapsed := time.Duration(0); elapsed < timeout; elapsed = elapsed + interval {
		if isDone() {
			return
		}
		time.Sleep(interval)
	}

	Fail("timed out")
}

func fooIsRunning() bool {
	return isRunning("foo")
}

func barIsRunning() bool {
	return isRunning("bar")
}

func bazIsRunning() bool {
	return isRunning("baz")
}

func isRunning(job string) bool {
	summary := getMonitSummary()
	pattern := getJobRunningPattern(job)
	return pattern.MatchString(summary)
}

func getJobRunningPattern(job string) *regexp.Regexp {
	pattern := fmt.Sprintf(`(?m)^Process '%s'\s+running$`, job)
	return regexp.MustCompile(pattern)
}

func startMonit() {
	cmd := exec.Command("monit", "-c", "/home/vcap/monitrc")
	err := cmd.Run()
	Expect(err).NotTo(HaveOccurred())
}

func getMonitSummary() string {
	cmd := exec.Command("monit", "-c", "/home/vcap/monitrc", "summary")
	summary, err := cmd.CombinedOutput()
	Expect(err).NotTo(HaveOccurred())
	return string(summary)
}
