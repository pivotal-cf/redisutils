package monit

import (
	"path/filepath"

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
			summary        []byte
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
			Expect(summary).To(Equal(exampleSummary))
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
	})
})

func getExampleMonitSummary() []byte {
	path := filepath.FromSlash("assets/example_monit_summary.txt")
	return readFile(path)
}
