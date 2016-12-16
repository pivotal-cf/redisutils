package monit_integration_test

import (
	"fmt"
	"os/exec"
	"regexp"
	"testing"

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
	Eventually(fooIsRunning, "15s").Should(BeTrue())
	Eventually(barIsRunning, "15s").Should(BeTrue())
	Eventually(bazIsRunning, "15s").Should(BeTrue())
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

func bazIsNotMonitored() bool {
	return isNotMonitored("baz")
}

func isNotMonitored(job string) bool {
	summary := getMonitSummary()
	pattern := getJobNotMonitoredPattern(job)
	return pattern.Match(summary)
}

func isRunning(job string) bool {
	summary := getMonitSummary()
	pattern := getJobRunningPattern(job)
	return pattern.Match(summary)
}

func getJobRunningPattern(job string) *regexp.Regexp {
	pattern := fmt.Sprintf(`(?m)^Process '%s'\s+running$`, job)
	return regexp.MustCompile(pattern)
}

func getJobNotMonitoredPattern(job string) *regexp.Regexp {
	pattern := fmt.Sprintf(`(?m)^Process '%s'\s+not monitored$`, job)
	return regexp.MustCompile(pattern)
}

func startMonit() {
	cmd := exec.Command("monit", "-c", "/home/vcap/monitrc")
	err := cmd.Run()
	Expect(err).NotTo(HaveOccurred())
}

func monitStartBaz() {
	cmd := exec.Command("monit", "-c", "/home/vcap/monitrc", "start", "baz")
	err := cmd.Run()
	Expect(err).NotTo(HaveOccurred())
}

func getMonitSummary() []byte {
	cmd := exec.Command("monit", "-c", "/home/vcap/monitrc", "summary")
	summary, err := cmd.CombinedOutput()
	Expect(err).NotTo(HaveOccurred())
	return summary
}
