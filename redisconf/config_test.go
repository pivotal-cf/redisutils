package redisconf

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("redisconf", func() {
	Describe("Config", func() {
		It("implements the Directive interface", func() {
			var _ Directive = Config{}
		})

		Describe("#String", func() {
			It("space separates Name and Value", func() {
				directive := NewConfig("save", "900 1")
				Expect(directive.String()).To(Equal("save 900 1"))
			})
		})

		Describe("#Validate", func() {
			It("does not return an error", func() {
				directive := NewConfig("save", "900 1")
				err := directive.Validate()
				Expect(err).NotTo(HaveOccurred())
			})

			Context("when Name is not a valid config", func() {
				validateErr := errors.New("unknown config: `foo`")

				It("returns an error", func() {
					directive := NewConfig("foo", "bar")
					err := directive.Validate()
					Expect(err).To(MatchError(validateErr))
				})
			})
		})

		Describe("#DecodeConfig", func() {
			It("does not return an error", func() {
				_, err := DecodeConfig("save 600 1")
				Expect(err).NotTo(HaveOccurred())
			})

			It("decodes the config", func() {
				config, _ := DecodeConfig("save 600 1")
				Expect(config).To(Equal(NewConfig("save", "600 1")))
			})

			Context("when config is not valid", func() {
				configErr := errors.New("unknown config: `foo`")

				It("returns an error", func() {
					_, err := DecodeConfig("foo bar")
					Expect(err).To(MatchError(configErr))
				})
			})

			Context("the the line is blank", func() {
				decodeErr := errors.New("unknown config: ``")

				It("returns an error", func() {
					_, err := DecodeConfig("\n")
					Expect(err).To(MatchError(decodeErr))
				})
			})
		})
	})
})
