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
	Eventually(allAreRunning, "15s").Should(BeTrue())
})

func fooIsRunning() bool {
	return isRunning("foo")
}

func barIsRunning() bool {
	return isRunning("bar")
}

func bazIsRunning() bool {
	return isRunning("baz")
}

func allAreRunning() bool {
	return fooIsRunning() &&
		barIsRunning() &&
		bazIsRunning()
}

func fooIsNotMonitored() bool {
	return isNotMonitored("baz")
}

func barIsNotMonitored() bool {
	return isNotMonitored("baz")
}

func bazIsNotMonitored() bool {
	return isNotMonitored("baz")
}

func allAreNotMonitored() bool {
	return fooIsNotMonitored() &&
		barIsNotMonitored() &&
		bazIsNotMonitored()
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

func monitStartAll() {
	cmd := exec.Command("monit", "-c", "/home/vcap/monitrc", "start", "all")
	err := cmd.Run()
	Expect(err).NotTo(HaveOccurred())
}

func monitStopBaz() {
	cmd := exec.Command("monit", "-c", "/home/vcap/monitrc", "stop", "baz")
	err := cmd.Run()
	Expect(err).NotTo(HaveOccurred())
}

func monitStopAll() {
	cmd := exec.Command("monit", "-c", "/home/vcap/monitrc", "stop", "all")
	err := cmd.Run()
	Expect(err).NotTo(HaveOccurred())
}

func getMonitSummary() []byte {
	cmd := exec.Command("monit", "-c", "/home/vcap/monitrc", "summary")
	summary, err := cmd.CombinedOutput()
	Expect(err).NotTo(HaveOccurred())
	return summary
}
