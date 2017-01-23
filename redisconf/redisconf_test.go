package redisconf_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pivotal-cf/redisutils/redisconf"
)

var _ = Describe("redisconf", func() {
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
