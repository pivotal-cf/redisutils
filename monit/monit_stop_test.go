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
			Expect(pureFake.Cmd.CombinedOutputCallCount()).To(Equal(1))
		})

		Context("when CombinedOutput returns an error", func() {
			combinedOutputErr := errors.New("Run failed")

			BeforeEach(func() {
				pureFake.Cmd.CombinedOutputReturns([]byte("Run failed"), errors.New(""))
			})

			It("it propagates the stdout message to the start error message", func() {
				Expect(stopErr).To(MatchError(combinedOutputErr))
			})
		})
	})

	Describe("#StopAndWait", func() {
		var (
			stopAndWaitErr error
			process        string
		)

		BeforeEach(func() {
			process = "foo"
			output := []byte("Process 'foo' not monitored")
			pureFake.Cmd.CombinedOutputReturns(output, nil)
		})

		JustBeforeEach(func() {
			stopAndWaitErr = monit.StopAndWait(process)
		})

		It("does not return an error", func() {
			Expect(stopAndWaitErr).NotTo(HaveOccurred())
		})

		It("calls `monit stop {job}`", func() {
			command := joinCommand(pureFake.Exec.CommandArgsForCall(0))
			Expect(command).To(Equal("monit stop foo"))
		})

		It("calls `monit summary`", func() {
			command := joinCommand(pureFake.Exec.CommandArgsForCall(1))
			Expect(command).To(Equal("monit summary"))
		})

		Context("when GetStatus returns StatusNotMoitored some time later", func() {
			BeforeEach(func() {
				pureFake.Cmd.CombinedOutputStub = combinedOutputReturns([]byteSliceAndError{
					{[]byte("Process 'foo' running"), nil},
					{[]byte("Process 'foo' running"), nil},
					{[]byte("Process 'foo' not monitored"), nil},
				}).sequentially

				monit.interval = time.Duration(0)
			})

			It("does not return an error", func() {
				Expect(stopAndWaitErr).NotTo(HaveOccurred())
			})

			It("calls `monit summary` multiple times", func() {
				Expect(pureFake.Cmd.CombinedOutputCallCount()).To(Equal(3))
			})
		})

		Context("when stopping all", func() {
			BeforeEach(func() {
				process = "all"
			})

			It("does not return an error", func() {
				Expect(stopAndWaitErr).NotTo(HaveOccurred())
			})

			It("calls `monit stop all`", func() {
				command := joinCommand(pureFake.Exec.CommandArgsForCall(0))
				Expect(command).To(Equal("monit stop all"))
			})

			It("calls `monit summary`", func() {
				command := joinCommand(pureFake.Exec.CommandArgsForCall(1))
				Expect(command).To(Equal("monit summary"))
			})

			Context("when GetSummary returns an error", func() {
				getSummaryError := errors.New("GetSummary failed")

				BeforeEach(func() {
					pureFake.Cmd.CombinedOutputStub = combinedOutputReturns([]byteSliceAndError{
						{nil, nil},
						{nil, getSummaryError},
					}).sequentially
				})

				It("returns the error", func() {
					Expect(stopAndWaitErr).To(MatchError(getSummaryError))
				})
			})

			Context("when GetSummary reports all not monitored some time later", func() {
				summaries := []byteSliceAndError{
					{[]byte("Process 'foo' running\nProcess 'bar' running"), nil},
					{[]byte("Process 'foo' not monitored\nProcess 'bar' running"), nil},
					{[]byte("Process 'foo' not monitored\nProcess 'bar' not monitored"), nil},
				}

				BeforeEach(func() {
					pureFake.Cmd.CombinedOutputStub = combinedOutputReturns(summaries).sequentially
					monit.interval = time.Duration(0)
				})

				It("does not return an error", func() {
					Expect(stopAndWaitErr).NotTo(HaveOccurred())
				})

				It("calls `monit summary` multiple times", func() {
					Expect(pureFake.Cmd.CombinedOutputCallCount()).To(Equal(3))
				})
			})
		})

		Context("when Stop returns an error", func() {
			stopErr := errors.New("Stop failed")

			BeforeEach(func() {
				pureFake.Cmd.CombinedOutputReturns([]byte("Stop failed"), errors.New(""))
			})

			It("returns the error", func() {
				Expect(stopAndWaitErr).To(MatchError(stopErr))
			})
		})

		Context("when GetStatus returns an error", func() {
			getStatusErr := errors.New("GetStatus failed")

			BeforeEach(func() {
				pureFake.Cmd.CombinedOutputReturns([]byte("GetStatus failed"), errors.New(""))
			})

			It("returns the error", func() {
				Expect(stopAndWaitErr).To(MatchError(getStatusErr))
			})
		})

		Context("when waiting times out", func() {
			BeforeEach(func() {
				monit.timeout = time.Duration(0)
				pureFake.Cmd.CombinedOutputReturns(nil, nil)
			})

			It("returns an error", func() {
				Expect(stopAndWaitErr).To(MatchError(ErrTimeout))
			})
		})
	})
})
