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

	Describe("GetSummary", func() {
		var (
			summary       map[string]int
			getSummaryErr error
		)

		BeforeEach(func() {
			summary, getSummaryErr = testMonit.GetSummary()
		})

		It("does not return an error", func() {
			Expect(getSummaryErr).NotTo(HaveOccurred())
		})

		It("reports processes running", func() {
			expectedSummary := map[string]int{
				"foo": monit.StatusRunning,
				"bar": monit.StatusRunning,
				"baz": monit.StatusRunning,
			}
			Expect(summary).To(Equal(expectedSummary))
		})
	})
})
