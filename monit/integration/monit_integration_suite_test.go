package monit_integration_test

import (
	"fmt"
	"os/exec"
	"regexp"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const monitrcPath = "/home/vcap/monitrc"

func TestMonitIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "monit integration suite")
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
	cmd := monitCommand()
	err := cmd.Run()
	Expect(err).NotTo(HaveOccurred())
}

func monitStart(job string) {
	cmd := monitCommand("start", job)
	err := cmd.Run()
	Expect(err).NotTo(HaveOccurred())
}

func monitStop(job string) {
	cmd := monitCommand("stop", job)
	err := cmd.Run()
	Expect(err).NotTo(HaveOccurred())
}

func getMonitSummary() []byte {
	cmd := monitCommand("summary")
	summary, err := cmd.CombinedOutput()
	Expect(err).NotTo(HaveOccurred())
	return summary
}

func monitCommand(args ...string) *exec.Cmd {
	cmd := exec.Command("monit", "-c", monitrcPath)
	cmd.Args = append(cmd.Args, args...)
	return cmd
}
