package redisconf

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("redisconf", func() {
	Describe("Conf", func() {
		It("assigns the keyword correctly", func() {
			var c Conf = NewConf("hello", "brother", "mine")
			Expect(c.Keyword).To(Equal("hello"))
		})

		It("assigns the args correctly", func() {
			var c Conf = NewConf("hello", "brother", "mine")
			expectedArgs := []string{"brother", "mine"}
			Expect(c.Args).To(Equal(expectedArgs))
		})

		Context("without args", func() {
			It("assigns the keyword correctly", func() {
				c := NewConf("hello")
				Expect(c.Keyword).To(Equal("hello"))
			})

			It("assigns the args correctly", func() {
				c := NewConf("hello")
				Expect(c.Args).To(BeEmpty())
			})
		})

		Describe("String", func() {
			It("converts to space separated string", func() {
				c := NewConf("hello", "brother", "mine")
				Expect(c.String()).To(Equal("hello brother mine"))
			})
		})
	})

	Describe("RedisConf", func() {
		It("exists", func() {
			var _ RedisConf = RedisConf{}
		})

		Describe("Encode", func() {
			It("encodes a RedisConf as newline separated Confs", func() {
				rc := RedisConf{
					NewConf("hello", "brother", "mine"),
					NewConf("did", "you", "miss", "me"),
				}
				Expect(rc.Encode()).To(Equal("hello brother mine\ndid you miss me\n"))
			})

			Context("when the RedisConf is empty", func() {
				rc := RedisConf{}

				It("returns an empty string", func() {
					Expect(rc.Encode()).To(Equal(""))
				})
			})
		})
	})

})
