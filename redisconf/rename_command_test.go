package redisconf

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("redisconf", func() {
	Describe("RenameCommand", func() {
		It("implements the Directive interface", func() {
			var _ Directive = RenameCommand{}
		})

		Describe("#String", func() {
			It("spaces separates Command and Alias after the rename directive", func() {
				directive := NewRenameCommand("BGSAVE", "foo")
				Expect(directive.String()).To(Equal("rename-command BGSAVE foo"))
			})

			Context("when Alias is blank", func() {
				It("returns empty quotes for its alias part", func() {
					directive := NewRenameCommand("BGSAVE", "")
					Expect(directive.String()).To(Equal(`rename-command BGSAVE ""`))
				})
			})
		})
	})
})
