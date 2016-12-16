package monit_integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pivotal-cf/redisutils/monit"
)

var _ = Describe("monit", func() {
	var testMonit *monit.Monit

	BeforeEach(func() {
		testMonit = monit.New()
		testMonit.MonitrcPath = "/home/vcap/monitrc"
	})

	Describe("#GetSummary", func() {
		var (
			summary       monit.Statuses
			getSummaryErr error
		)

		BeforeEach(func() {
			summary, getSummaryErr = testMonit.GetSummary()
		})

		It("does not return an error", func() {
			Expect(getSummaryErr).NotTo(HaveOccurred())
		})

		It("reports processes running", func() {
			expectedSummary := monit.Statuses{
				"foo": monit.StatusRunning,
				"bar": monit.StatusRunning,
				"baz": monit.StatusRunning,
			}
			Expect(summary).To(Equal(expectedSummary))
		})
	})

	Describe("#GetStatus", func() {
		var (
			status       monit.Status
			getStatusErr error
		)

		BeforeEach(func() {
			status, getStatusErr = testMonit.GetStatus("baz")
		})

		It("does not return an error", func() {
			Expect(getStatusErr).NotTo(HaveOccurred())
		})

		It("gets the correct status", func() {
			Expect(status).To(Equal(monit.StatusRunning))
		})
	})

	Describe("#Stop", func() {
		var stopErr error

		BeforeEach(func() {
			stopErr = testMonit.Stop("baz")
		})

		AfterEach(func() {
			Eventually(bazIsNotMonitored, "10s").Should(BeTrue())
			monitStartBaz()
			Eventually(bazIsRunning, "15s").Should(BeTrue())
		})

		It("stops baz", func() {
			By("not returning and error")
			Expect(stopErr).NotTo(HaveOccurred())

			By("and stopping baz")
			Eventually(bazIsNotMonitored, "10s").Should(BeTrue())
		})
	})
})
