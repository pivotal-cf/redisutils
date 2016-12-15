package monit_integration_test

import (
	"os/exec"
	"regexp"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMonitIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Monit Integration Suite")
}

var _ = BeforeSuite(func() {
	startMonit()
	waitForFooRunning()
})

func waitForFooRunning() {
	interval := time.Millisecond * 200
	timeout := time.Second * 15

	for elapsed := time.Duration(0); elapsed < timeout; elapsed = elapsed + interval {
		if fooIsRunning() {
			return
		}
		time.Sleep(interval)
	}

	Fail("timed out waiting for monit to start")
}

func fooIsRunning() bool {
	summary := getMonitSummary()
	pattern := regexp.MustCompile(`(?m)^Process 'foo'\s+running$`)
	return pattern.MatchString(summary)
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
