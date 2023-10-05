package redisconf

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("redisconf", func() {
	Describe("#New", func() {
		It("does not return an error", func() {
			_, err := New()
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when given an invalid directive", func() {
			validateErr := errors.New("unknown config: `foo`")

			It("returns an error", func() {
				_, err := New(NewConfig("foo", "bar"))
				Expect(err).To(MatchError(validateErr))
			})
		})
	})

	Describe("#String", func() {
		It("returns an empty string", func() {
			redisConf, _ := New()
			Expect(redisConf.String()).To(Equal(""))
		})

		Context("when it contains one directive", func() {
			It("newline returns that directive as a line", func() {
				redisConf, _ := New(NewConfig("save", "600 1"))
				Expect(redisConf.String()).To(Equal("save 600 1\n"))
			})
		})

		Context("when it contains many lines", func() {
			It("newline separates each directive", func() {
				expected := "save 600 1\nrename-command BGSAVE \"\"\ndaemonize yes\n"

				redisConf, _ := New(
					NewConfig("save", "600 1"),
					NewRenameCommand("BGSAVE", ""),
					NewConfig("daemonize", "yes"),
				)
				Expect(redisConf.String()).To(Equal(expected))
			})
		})
	})
})
