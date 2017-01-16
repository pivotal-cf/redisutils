package redisconf

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("redisconf", func() {
	var c Conf

	Describe("Conf", func() {
		BeforeEach(func() {
			c = NewConf("hello", "brother", "mine")
		})

		It("assigns the keyword correctly", func() {
			Expect(c.Keyword).To(Equal("hello"))
		})

		It("assigns the args correctly", func() {
			expectedArgs := []string{"brother", "mine"}
			Expect(c.Args).To(Equal(expectedArgs))
		})

		Context("without args", func() {
			BeforeEach(func() {
				c = NewConf("hello")
			})

			It("assigns the keyword correctly", func() {
				Expect(c.Keyword).To(Equal("hello"))
			})

			It("assigns the args correctly", func() {
				Expect(c.Args).To(BeEmpty())
			})
		})

		Describe("String", func() {
			It("converts to space separated string", func() {
				Expect(c.String()).To(Equal("hello brother mine"))
			})
		})
	})

	Describe("RedisConf", func() {
		Describe("Encode", func() {
			var redisConf RedisConf

			BeforeEach(func() {
				redisConf = RedisConf{
					NewConf("hello", "brother", "mine"),
					NewConf("did", "you", "miss", "me"),
				}
			})

			It("encodes a RedisConf as newline separated Confs", func() {
				Expect(redisConf.Encode()).To(Equal("hello brother mine\ndid you miss me\n"))
			})

			Context("when the RedisConf is empty", func() {
				BeforeEach(func() {
					redisConf = RedisConf{}
				})

				It("returns an empty string", func() {
					Expect(redisConf.Encode()).To(Equal(""))
				})
			})
		})

		Describe("Append", func() {
			var (
				redisConf   RedisConf
				appendError error
			)

			BeforeEach(func() {
				redisConf = RedisConf{NewConf("hello", "brother", "mine")}
				appendError = redisConf.Append(NewConf("hi", "mycroft"))
			})

			It("succeeds", func() {
				Expect(appendError).NotTo(HaveOccurred())
			})

			It("appends the Conf", func() {
				expectedConf := RedisConf{
					NewConf("hello", "brother", "mine"),
					NewConf("hi", "mycroft"),
				}
				Expect(redisConf).To(Equal(expectedConf))
			})
		})
	})

})
