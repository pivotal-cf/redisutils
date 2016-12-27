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

	Describe("#StopAndWait", func() {
		var stopAndWaitErr error

		BeforeEach(func() {
			output := []byte("Process 'foo' not monitored")
			pureFake.Cmd.CombinedOutputReturns(output, nil)
		})

		JustBeforeEach(func() {
			stopAndWaitErr = monit.StopAndWait("foo")
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
				pureFake.Cmd.CombinedOutputStub = combinedOutputReturns([][]byte{
					[]byte("Process 'foo' running"),
					[]byte("Process 'foo' running"),
					[]byte("Process 'foo' not monitored"),
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

		Context("when Stop returns an error", func() {
			stopErr := errors.New("Stop failed")

			BeforeEach(func() {
				pureFake.Cmd.RunReturns(stopErr)
			})

			It("returns the error", func() {
				Expect(stopAndWaitErr).To(MatchError(stopErr))
			})
		})

		Context("when GetStatus returns an error", func() {
			getStatusErr := errors.New("GetStatus failed")

			BeforeEach(func() {
				pureFake.Cmd.CombinedOutputReturns(nil, getStatusErr)
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
