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
		monit *SysMonit
		fakes *iexec.NestedCommandFake
	)

	BeforeEach(func() {
		monit = New()
		fakes = iexec.NewNestedCommandFake()
		monit.exec = fakes.Exec
	})

	Describe("#Start", func() {
		var (
			startErr error
			command  string
			args     []string
		)

		JustBeforeEach(func() {
			startErr = monit.Start("foo")
			Expect(fakes.Exec.CommandCallCount()).To(Equal(1))
			command, args = fakes.Exec.CommandArgsForCall(0)
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
			Expect(fakes.Cmd.CombinedOutputCallCount()).To(Equal(1))
		})

		Context("when CombinedOutput returns an error", func() {
			runErr := errors.New("CombinedOutput failed")

			BeforeEach(func() {
				fakes.Cmd.CombinedOutputReturns([]byte("CombinedOutput failed"), errors.New(""))
			})

			It("it propagates the stdout message to the start error message", func() {
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
			fakes.Cmd.CombinedOutputReturns(output, nil)
		})

		JustBeforeEach(func() {
			startAndWaitErr = monit.StartAndWait(process)
		})

		It("does not return an error", func() {
			Expect(startAndWaitErr).NotTo(HaveOccurred())
		})

		It("calls `monit stop {job}`", func() {
			command := joinCommand(fakes.Exec.CommandArgsForCall(0))
			Expect(command).To(Equal("monit start foo"))
		})

		It("calls `monit summary`", func() {
			command := joinCommand(fakes.Exec.CommandArgsForCall(1))
			Expect(command).To(Equal("monit summary"))
		})

		Context("When GetStatus returns StatusRunning some time later", func() {
			BeforeEach(func() {
				fakes.Cmd.CombinedOutputStub = combinedOutputReturns([]byteSliceAndError{
					{[]byte("Process 'foo' not monitored"), nil},
					{[]byte("Process 'foo' not monitored"), nil},
					{[]byte("Process 'foo' running"), nil},
				}).sequentially

				monit.interval = time.Duration(0)
			})

			It("does not return an error", func() {
				Expect(startAndWaitErr).NotTo(HaveOccurred())
			})

			It("calls `monit summary` multiple times", func() {
				Expect(fakes.Cmd.CombinedOutputCallCount()).To(Equal(3))
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
				command := joinCommand(fakes.Exec.CommandArgsForCall(0))
				Expect(command).To(Equal("monit start all"))
			})

			It("calls `monit summary`", func() {
				command := joinCommand(fakes.Exec.CommandArgsForCall(1))
				Expect(command).To(Equal("monit summary"))
			})

			Context("when GetSummary returns an error", func() {
				getSummaryError := errors.New("GetSummary failed")

				BeforeEach(func() {
					fakes.Cmd.CombinedOutputStub = combinedOutputReturns([]byteSliceAndError{
						{nil, nil},
						{nil, getSummaryError},
					}).sequentially
				})

				It("returns the error", func() {
					Expect(startAndWaitErr).To(MatchError(getSummaryError))
				})
			})

			Context("when GetSummary reports all not monitored some time later", func() {
				summaries := []byteSliceAndError{
					{[]byte("Process 'foo' not monitored\nProcess 'bar' not monitored"), nil},
					{[]byte("Process 'foo' not monitored\nProcess 'bar' running"), nil},
					{[]byte("Process 'foo' running\nProcess 'bar' running"), nil},
				}

				BeforeEach(func() {
					fakes.Cmd.CombinedOutputStub = combinedOutputReturns(summaries).sequentially
					monit.interval = time.Duration(0)
				})

				It("does not return an error", func() {
					Expect(startAndWaitErr).NotTo(HaveOccurred())
				})

				It("calls `monit summary` multiple times", func() {
					Expect(fakes.Cmd.CombinedOutputCallCount()).To(Equal(3))
				})
			})
		})

		Context("when Start returns an error", func() {
			startErr := errors.New("Start failed")

			BeforeEach(func() {
				fakes.Cmd.CombinedOutputReturns([]byte("Start failed"), startErr)
			})

			It("returns the error", func() {
				Expect(startAndWaitErr).To(MatchError(startErr))
			})
		})

		Context("when GetStatus returns an error", func() {
			getStatusErr := errors.New("GetStatus failed")

			BeforeEach(func() {
				fakes.Cmd.CombinedOutputReturns([]byte("GetStatus failed"), getStatusErr)
			})

			It("returns the error", func() {
				Expect(startAndWaitErr).To(MatchError(getStatusErr))
			})
		})

		Context("when waiting times out", func() {
			BeforeEach(func() {
				monit.timeout = time.Duration(0)
				fakes.Cmd.CombinedOutputReturns(nil, nil)
			})

			It("returns an error", func() {
				Expect(startAndWaitErr).To(MatchError(ErrTimeout))
			})
		})
	})
})
