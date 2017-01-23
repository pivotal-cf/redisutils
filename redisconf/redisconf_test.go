package redisconf_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pivotal-cf/redisutils/redisconf"
)

var _ = Describe("redisconf", func() {
	Describe("Config", func() {
		It("implements the Directive interface", func() {
			var _ redisconf.Directive = redisconf.Config{}
		})

		Describe("#String", func() {
			It("space separates Name and Value", func() {
				directive := redisconf.NewConfig("save", "900 1")
				Expect(directive.String()).To(Equal("save 900 1"))
			})
		})

		Describe("#Validate", func() {
			It("does not return an error", func() {
				directive := redisconf.NewConfig("save", "900 1")
				err := directive.Validate()
				Expect(err).NotTo(HaveOccurred())
			})

			Context("when Name is not a valid config", func() {
				validateErr := errors.New("unknown config: foo")

				It("returns an error", func() {
					directive := redisconf.NewConfig("foo", "bar")
					err := directive.Validate()
					Expect(err).To(MatchError(validateErr))
				})
			})
		})
	})

	Describe("RenameCommand", func() {
		It("implements the Directive interface", func() {
			var _ redisconf.Directive = redisconf.RenameCommand{}
		})

		Describe("#String", func() {
			It("spaces separates Command and Alias after the rename directive", func() {
				directive := redisconf.NewRenameCommand("BGSAVE", "foo")
				Expect(directive.String()).To(Equal("rename-command BGSAVE foo"))
			})

			Context("when Alias is blank", func() {
				It("returns empty quotes for its alias part", func() {
					directive := redisconf.NewRenameCommand("BGSAVE", "")
					Expect(directive.String()).To(Equal(`rename-command BGSAVE ""`))
				})
			})
		})
	})
})
