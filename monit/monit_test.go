package monit

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/BooleanCat/igo/ios/iexec"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("monit", func() {
	var monit *Monit
	var pureFake *iexec.PureFake

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

	Describe("#Stop", func() {
		var (
			stopErr error
			command string
			args    []string
		)

		JustBeforeEach(func() {
			stopErr = monit.Stop("foo")
			Expect(pureFake.Exec.CommandCallCount()).To(Equal(1))
			command, args = pureFake.Exec.CommandArgsForCall(0)
		})

		It("does not return an error", func() {
			Expect(stopErr).NotTo(HaveOccurred())
		})

		It("prepares to execute the monit binary", func() {
			Expect(command).To(Equal("monit"))
		})

		It("prepares to execute `monit stop foo`", func() {
			Expect(args).To(Equal([]string{"stop", "foo"}))
		})

		It("runs `monit stop foo`", func() {
			Expect(pureFake.Cmd.RunCallCount()).To(Equal(1))
		})

		Context("when run returns an error", func() {
			runErr := errors.New("Run failed")

			BeforeEach(func() {
				pureFake.Cmd.RunReturns(runErr)
			})

			It("returns the error", func() {
				Expect(stopErr).To(MatchError(runErr))
			})
		})
	})
})

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
