package monit_integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pivotal-cf/redisutils/monit"
)

var _ = Describe("monit", func() {
	var testMonit monit.Monit

	BeforeEach(func() {
		testMonit = monit.New()
		testMonit.SetMonitrcPath("/home/vcap/monitrc")
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

	Describe("#Start", func() {
		var startErr error

		BeforeEach(func() {
			monitStopBaz()
			Eventually(bazIsNotMonitored, "15s").Should(BeTrue())
			startErr = testMonit.Start("baz")
		})

		It("starts baz", func() {
			By("not returning an error")
			Expect(startErr).NotTo(HaveOccurred())

			By("and starting baz")
			Eventually(bazIsRunning, "10s").Should(BeTrue())
		})
	})

	Describe("#StartAndWait", func() {
		var startAndWaitErr error

		BeforeEach(func() {
			monitStopBaz()
			Eventually(bazIsNotMonitored, "15s").Should(BeTrue())
			startAndWaitErr = testMonit.StartAndWait("baz")
		})

		It("starts baz", func() {
			By("not returning an error")
			Expect(startAndWaitErr).NotTo(HaveOccurred())

			By("and starting baz")
			Expect(bazIsRunning()).To(BeTrue())
		})
	})

	Describe("#Stop", func() {
		var stopErr error

		BeforeEach(func() {
			stopErr = testMonit.Stop("baz")
		})

		AfterEach(func() {
			Eventually(bazIsNotMonitored, "15s").Should(BeTrue())
			monitStartBaz()
			Eventually(bazIsRunning, "15s").Should(BeTrue())
		})

		It("stops baz", func() {
			By("not returning an error")
			Expect(stopErr).NotTo(HaveOccurred())

			By("and stopping baz")
			Eventually(bazIsNotMonitored, "15s").Should(BeTrue())
		})
	})

	Describe("#StopAndWait", func() {
		var stopAndWaitErr error

		BeforeEach(func() {
			stopAndWaitErr = testMonit.StopAndWait("baz")
		})

		AfterEach(func() {
			Eventually(bazIsNotMonitored, "15s").Should(BeTrue())
			monitStartBaz()
			Eventually(bazIsRunning, "15s").Should(BeTrue())
		})

		It("stops baz", func() {
			By("not returning an error")
			Expect(stopAndWaitErr).NotTo(HaveOccurred())

			By("and stopping baz")
			Expect(bazIsNotMonitored()).To(BeTrue())
		})
	})

	It("can stop and start all processes", func() {
		By("stopping all processes")
		err := testMonit.Stop("all")
		Expect(err).NotTo(HaveOccurred())
		Eventually(allAreNotMonitored, "15s").Should(BeTrue())

		By("starting all processes")
		err = testMonit.Start("all")
		Expect(err).NotTo(HaveOccurred())
		Eventually(allAreRunning, "15s").Should(BeTrue())
	})
})
