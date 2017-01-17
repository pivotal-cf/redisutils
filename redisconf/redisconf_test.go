package redisconf_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pivotal-cf/redisutils/redisconf"
)

var _ = Describe("redisconf", func() {
	Describe("Directive", func() {
		var directive redisconf.Directive

		Describe("#NewDirective", func() {
			BeforeEach(func() {
				directive = redisconf.NewDirective("hello", "brother", "mine")
			})

			It("assigns Keyword correctly", func() {
				Expect(directive.Keyword).To(Equal("hello"))
			})

			It("assigns Args correctly", func() {
				expectedArgs := []string{"brother", "mine"}
				Expect(directive.Args).To(Equal(expectedArgs))
			})

			Context("when no Args are provided", func() {
				BeforeEach(func() {
					directive = redisconf.NewDirective("hello")
				})

				It("assigns Keyword correctly", func() {
					Expect(directive.Keyword).To(Equal("hello"))
				})

				It("assigns Args correctly", func() {
					Expect(directive.Args).To(BeEmpty())
				})
			})
		})

		Describe("String", func() {
			BeforeEach(func() {
				directive = redisconf.NewDirective("hello", "brother", "mine")
			})

			It("converts to space separated string", func() {
				Expect(directive.String()).To(Equal("hello brother mine"))
			})
		})
	})

	Describe("RedisConf", func() {
		var redisConf redisconf.RedisConf

		Describe("Encode", func() {
			BeforeEach(func() {
				redisConf = redisconf.New()
			})

			It("returns an empty string", func() {
				Expect(redisConf.Encode()).To(Equal(""))
			})

			Context("when the RedisConf is not empty", func() {
				BeforeEach(func() {
					redisConf = redisconf.New(
						redisconf.NewDirective("hello", "brother", "mine"),
						redisconf.NewDirective("did", "you", "miss", "me"),
					)
				})

				It("returns newline separated Directives", func() {
					Expect(redisConf.Encode()).To(Equal("hello brother mine\ndid you miss me\n"))
				})
			})
		})

		Describe("Append", func() {
			directive := redisconf.NewDirective("foo", "bar")

			BeforeEach(func() {
				redisConf = redisconf.New()
				redisConf = redisConf.Append(directive)
			})

			It("appends the directive", func() {
				expectedConf := redisconf.New(directive)
				Expect(redisConf).To(Equal(expectedConf))
			})

			Context("when a directive already exists", func() {
				BeforeEach(func() {
					redisConf = redisConf.Append(directive)
				})

				It("does nothing", func() {
					expectedConf := redisconf.New(directive)
					Expect(redisConf).To(Equal(expectedConf))
				})
			})
		})
	})
})
