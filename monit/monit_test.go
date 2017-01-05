package monit

import (
	"errors"
	"strings"

	"github.com/BooleanCat/igo/ios/iexec"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("monit", func() {
	var (
		monit    *SysMonit
		pureFake *iexec.PureFake
	)

	BeforeEach(func() {
		monit = New()
		pureFake = iexec.NewPureFake()
		monit.exec = pureFake.Exec
	})

	Describe("#GetSummary", func() {
		var (
			summary        Statuses
			getSummaryErr  error
			command        string
			args           []string
			exampleSummary []byte
		)

		BeforeEach(func() {
			exampleSummary = getExampleMonitSummary()
			pureFake.Cmd.CombinedOutputReturns(exampleSummary, nil)
		})

		JustBeforeEach(func() {
			summary, getSummaryErr = monit.GetSummary()
			Expect(pureFake.Exec.CommandCallCount()).To(Equal(1))
			command, args = pureFake.Exec.CommandArgsForCall(0)
		})

		It("does not return an error", func() {
			Expect(getSummaryErr).NotTo(HaveOccurred())
		})

		It("returns the output of `monit summary`", func() {
			expectedSummary := Statuses{
				"process-watcher":   StatusRunning,
				"process-destroyer": StatusRunning,
				"cf-redis-broker":   StatusRunning,
				"broker-nginx":      StatusRunning,
				"route_registrar":   StatusRunning,
			}
			Expect(summary).To(Equal(expectedSummary))
		})

		It("prepares to execute the monit binary", func() {
			Expect(command).To(Equal("monit"))
		})

		It("prepares to execute `monit summary`", func() {
			Expect(args).To(Equal([]string{"summary"}))
		})

		It("runs `monit summary`", func() {
			Expect(pureFake.Cmd.CombinedOutputCallCount()).To(Equal(1))
		})

		It("doesn't specify a monitrc path when calling `monit summary`", func() {
			joinedArgs := strings.Join(args, " ")
			Expect(joinedArgs).NotTo(ContainSubstring("-c"))
		})

		Context("when combined output returns an error", func() {
			combinedOutputErr := errors.New("CombinedOutput failed")

			BeforeEach(func() {
				pureFake.Cmd.CombinedOutputReturns(nil, combinedOutputErr)
			})

			It("returns the error", func() {
				Expect(getSummaryErr).To(MatchError(combinedOutputErr))
			})
		})

		Context("when a monitrc path is configured", func() {
			BeforeEach(func() {
				monit.MonitrcPath = "/foo/bar"
			})

			It("specifies that path when calling `monit summary`", func() {
				joinedArgs := strings.Join(args, " ")
				Expect(joinedArgs).To(ContainSubstring("-c " + monit.MonitrcPath))
			})
		})

		Context("when a monit executable is configured", func() {
			BeforeEach(func() {
				monit.Executable = "/foo/bar"
			})

			It("runs that executable", func() {
				Expect(command).To(Equal("/foo/bar"))
			})
		})

		Context("when one process is `stopped`", func() {
			expectedSummary := Statuses{
				"process-watcher":   StatusRunning,
				"process-destroyer": StatusNotMonitored,
				"cf-redis-broker":   StatusRunning,
				"broker-nginx":      StatusRunning,
				"route_registrar":   StatusRunning,
			}

			BeforeEach(func() {
				exampleSummary = getExampleMonitSummaryOneStopped()
				pureFake.Cmd.CombinedOutputReturns(exampleSummary, nil)
			})

			It("does not return an error", func() {
				Expect(getSummaryErr).NotTo(HaveOccurred())
			})

			It("reports the correct process as stopped", func() {
				Expect(summary).To(Equal(expectedSummary))
			})
		})

		Context("for all possible statuses", func() {
			expectedSummary := Statuses{
				"process-watcher":   StatusRunning,
				"process-destroyer": StatusNotMonitored,
				"cf-redis-broker":   StatusNotMonitoredStartPending,
				"broker-nginx":      StatusDoesNotExist,
				"route_registrar":   StatusInitializing,
				"crazy-job":         StatusNotMonitoredStopPending,
				"crazy-job-2":       StatusRunningRestartPending,
			}

			BeforeEach(func() {
				exampleSummary = getExampleMonitSummaryAllStatuses()
				pureFake.Cmd.CombinedOutputReturns(exampleSummary, nil)
			})

			It("does not return an error", func() {
				Expect(getSummaryErr).NotTo(HaveOccurred())
			})

			It("reports the correct status", func() {
				Expect(summary).To(Equal(expectedSummary))
			})
		})
	})

	Describe("#GetStatus", func() {
		var (
			status       Status
			getStatusErr error
			job          string
		)

		BeforeEach(func() {
			job = "broker-nginx"
			exampleSummary := getExampleMonitSummaryAllStatuses()
			pureFake.Cmd.CombinedOutputReturns(exampleSummary, nil)
		})

		JustBeforeEach(func() {
			status, getStatusErr = monit.GetStatus(job)
		})

		It("does not return an error", func() {
			Expect(getStatusErr).NotTo(HaveOccurred())
		})

		It("gets the correct status", func() {
			Expect(status).To(Equal(StatusDoesNotExist))
		})

		Context("when GetSummary returns an error", func() {
			combinedOutputErr := errors.New("CombinedOutput failed")

			BeforeEach(func() {
				pureFake.Cmd.CombinedOutputReturns(nil, combinedOutputErr)
			})

			It("returns the error", func() {
				Expect(getStatusErr).To(MatchError(combinedOutputErr))
			})
		})

		Context("when job doesn't exist", func() {
			noSuchJobErr := errors.New("no such job: `bar`")

			BeforeEach(func() {
				job = "bar"
			})

			It("returns an error", func() {
				Expect(getStatusErr).To(MatchError(noSuchJobErr))
			})
		})
	})

	Describe("#SetMonitrcPath", func() {
		It("sets correctly", func() {
			monit.SetMonitrcPath("/foo/bar")
			Expect(monit.MonitrcPath).To(Equal("/foo/bar"))
		})
	})

	Describe("#SetExecutable", func() {
		It("sets correctly", func() {
			monit.SetExecutable("foo/bar")
			Expect(monit.Executable).To(Equal("foo/bar"))
		})
	})
})
