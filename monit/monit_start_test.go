package monit

import (
	"errors"
	"time"

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

	Describe("#Start", func() {
		var (
			startErr error
			command  string
			args     []string
		)

		JustBeforeEach(func() {
			startErr = monit.Start("foo")
			Expect(pureFake.Exec.CommandCallCount()).To(Equal(1))
			command, args = pureFake.Exec.CommandArgsForCall(0)
		})

		It("does not return an error", func() {
			Expect(startErr).NotTo(HaveOccurred())
		})

		It("prepares to execute the monit binary", func() {
			Expect(command).To(Equal("monit"))
		})

		It("prepares to execute `monit start foo`", func() {
			Expect(args).To(Equal([]string{"start", "foo"}))
		})

		It("runs `monit start foo`", func() {
			Expect(pureFake.Cmd.RunCallCount()).To(Equal(1))
		})

		Context("when run returns an error", func() {
			runErr := errors.New("Run failed")

			BeforeEach(func() {
				pureFake.Cmd.RunReturns(runErr)
			})

			It("returns the error", func() {
				Expect(startErr).To(MatchError(runErr))
			})
		})
	})

	Describe("#StartAndWait", func() {
		var (
			startAndWaitErr error
			process         string
		)

		BeforeEach(func() {
			process = "foo"
			output := []byte("Process 'foo' running")
			pureFake.Cmd.CombinedOutputReturns(output, nil)
		})

		JustBeforeEach(func() {
			startAndWaitErr = monit.StartAndWait(process)
		})

		It("does not return an error", func() {
			Expect(startAndWaitErr).NotTo(HaveOccurred())
		})

		It("calls `monit stop {job}`", func() {
			command := joinCommand(pureFake.Exec.CommandArgsForCall(0))
			Expect(command).To(Equal("monit start foo"))
		})

		It("calls `monit summary`", func() {
			command := joinCommand(pureFake.Exec.CommandArgsForCall(1))
			Expect(command).To(Equal("monit summary"))
		})

		Context("When GetStatus returns StatusRunning some time later", func() {
			BeforeEach(func() {
				pureFake.Cmd.CombinedOutputStub = combinedOutputReturns([][]byte{
					[]byte("Process 'foo' not monitored"),
					[]byte("Process 'foo' not monitored"),
					[]byte("Process 'foo' running"),
				}).sequentially

				monit.interval = time.Duration(0)
			})

			It("does not return an error", func() {
				Expect(startAndWaitErr).NotTo(HaveOccurred())
			})

			It("calls `monit summary` multiple times", func() {
				Expect(pureFake.Cmd.CombinedOutputCallCount()).To(Equal(3))
			})
		})

		Context("when starting all", func() {
			BeforeEach(func() {
				process = "all"
			})

			It("does not return an error", func() {
				Expect(startAndWaitErr).NotTo(HaveOccurred())
			})

			It("calls `monit start all`", func() {
				command := joinCommand(pureFake.Exec.CommandArgsForCall(0))
				Expect(command).To(Equal("monit start all"))
			})

			It("calls `monit summary`", func() {
				command := joinCommand(pureFake.Exec.CommandArgsForCall(1))
				Expect(command).To(Equal("monit summary"))
			})

			Context("when GetSummary returns an error", func() {
				getSummaryError := errors.New("GetSummary failed")

				BeforeEach(func() {
					pureFake.Cmd.CombinedOutputReturns(nil, getSummaryError)
				})

				It("returns the error", func() {
					Expect(startAndWaitErr).To(MatchError(getSummaryError))
				})
			})

			Context("when GetSummary reports all not monitored some time later", func() {
				summaries := [][]byte{
					[]byte("Process 'foo' not monitored\nProcess 'bar' not monitored"),
					[]byte("Process 'foo' not monitored\nProcess 'bar' running"),
					[]byte("Process 'foo' running\nProcess 'bar' running"),
				}

				BeforeEach(func() {
					pureFake.Cmd.CombinedOutputStub = combinedOutputReturns(summaries).sequentially
					monit.interval = time.Duration(0)
				})

				It("does not return an error", func() {
					Expect(startAndWaitErr).NotTo(HaveOccurred())
				})

				It("calls `monit summary` multiple times", func() {
					Expect(pureFake.Cmd.CombinedOutputCallCount()).To(Equal(3))
				})
			})
		})

		Context("when Start returns an error", func() {
			startErr := errors.New("Start failed")

			BeforeEach(func() {
				pureFake.Cmd.RunReturns(startErr)
			})

			It("returns the error", func() {
				Expect(startAndWaitErr).To(MatchError(startErr))
			})
		})

		Context("when GetStatus returns an error", func() {
			getStatusErr := errors.New("GetStatus failed")

			BeforeEach(func() {
				pureFake.Cmd.CombinedOutputReturns(nil, getStatusErr)
			})

			It("returns the error", func() {
				Expect(startAndWaitErr).To(MatchError(getStatusErr))
			})
		})

		Context("when waiting times out", func() {
			BeforeEach(func() {
				monit.timeout = time.Duration(0)
				pureFake.Cmd.CombinedOutputReturns(nil, nil)
			})

			It("returns an error", func() {
				Expect(startAndWaitErr).To(MatchError(ErrTimeout))
			})
		})
	})
})
