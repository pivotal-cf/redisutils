package redisconf_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pivotal-cf/redisutils/redisconf"
)

var _ = Describe("redisconf", func() {
	Describe("Directive", func() {
		var (
			directive    redisconf.Directive
			directiveErr error
		)

		Describe("#NewDirective", func() {
			BeforeEach(func() {
				directive, directiveErr = redisconf.NewDirective("hello", "brother", "mine")
			})

			It("does not return an error", func() {
				Expect(directiveErr).NotTo(HaveOccurred())
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
					directive, directiveErr = redisconf.NewDirective("hello")
				})

				It("does not return an error", func() {
					Expect(directiveErr).NotTo(HaveOccurred())
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
				directive = newDirective("hello", "brother", "mine")
			})

			It("converts to space separated string", func() {
				Expect(directive.String()).To(Equal("hello brother mine"))
			})
		})
	})

	Describe("RedisConf", func() {
		var (
			redisConf    redisconf.RedisConf
			redisConfErr error
		)

		Describe("#New", func() {
			BeforeEach(func() {
				redisConf, redisConfErr = redisconf.New()
			})

			It("does not return an error", func() {
				Expect(redisConfErr).NotTo(HaveOccurred())
			})

			It("initialises empty", func() {
				Expect(redisConf).To(HaveLen(0))
			})

			Context("when given a directive", func() {
				BeforeEach(func() {
					redisConf, redisConfErr = redisconf.New(
						redisconf.Directive{"foo", redisconf.Args{"bar"}},
					)
				})

				It("does not return an error", func() {
					Expect(redisConfErr).NotTo(HaveOccurred())
				})

				It("has length 1", func() {
					Expect(redisConf).To(HaveLen(1))
				})

				It("contains the directive", func() {
					expectedDirective := redisconf.Directive{"foo", redisconf.Args{"bar"}}
					Expect(redisConf).To(ContainElement(expectedDirective))
				})
			})

			Context("when given 3 directives", func() {
				BeforeEach(func() {
					directives := []redisconf.Directive{
						{"foo", redisconf.Args{"bar"}},
						{"bar", redisconf.Args{"baz"}},
						{"baz", redisconf.Args{"boo", "baa"}},
					}
					redisConf, redisConfErr = redisconf.New(directives...)
				})

				It("does not return an error", func() {
					Expect(redisConfErr).NotTo(HaveOccurred())
				})

				It("has length 3", func() {
					Expect(redisConf).To(HaveLen(3))
				})

				It("contains the directives", func() {
					Expect(redisConf).To(ContainElement(redisconf.Directive{"foo", redisconf.Args{"bar"}}))
					Expect(redisConf).To(ContainElement(redisconf.Directive{"bar", redisconf.Args{"baz"}}))
					Expect(redisConf).To(ContainElement(redisconf.Directive{"baz", redisconf.Args{"boo", "baa"}}))
				})
			})
		})

		Describe("Encode", func() {
			BeforeEach(func() {
				redisConf = newRedisConf()
			})

			It("returns an empty string", func() {
				Expect(redisConf.Encode()).To(Equal(""))
			})

			Context("when the RedisConf is not empty", func() {
				BeforeEach(func() {
					directives := []redisconf.Directive{
						{"hello", redisconf.Args{"brother", "mine"}},
						{"did", redisconf.Args{"you", "miss", "me"}},
					}
					redisConf = newRedisConf(directives...)
				})

				It("returns newline separated Directives", func() {
					Expect(redisConf.Encode()).To(Equal("hello brother mine\ndid you miss me\n"))
				})
			})
		})

		Describe("Append", func() {
			BeforeEach(func() {
				directive := newDirective("foo", "bar")
				redisConf = newRedisConf()
				redisConf = redisConf.Append(directive)
			})

			It("appends the directive", func() {
				expectedConf := newRedisConf(newDirective("foo", "bar"))
				Expect(redisConf).To(Equal(expectedConf))
			})
		})
	})

})
