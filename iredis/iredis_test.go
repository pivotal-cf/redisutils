package iredis_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/pivotal-cf/redisutils/iredis"
)

var _ = Describe("interfaces", func() {
	Describe("Redis", func() {
		It("is implemented by iredis.RedisWrap", func() {
			var _ iredis.Redis = new(iredis.RedisWrap)
		})

		It("is implemented by iredis.RedisFake", func() {
			var _ iredis.Redis = new(iredis.RedisFake)
		})
	})

	Describe("Client", func() {
		It("is implemented by iredis.ClientWrap", func() {
			var _ iredis.Client = new(iredis.ClientWrap)
		})

		It("is implemented by iredis.ClientFake", func() {
			var _ iredis.Client = new(iredis.ClientFake)
		})
	})

	Describe("StatusCmd", func() {
		It("is implemented by iredis.StatusCmdWrap", func() {
			var _ iredis.StatusCmd = new(iredis.StatusCmdWrap)
		})

		It("is implemented by iredis.StatusCmdFake", func() {
			var _ iredis.StatusCmd = new(iredis.StatusCmdFake)
		})
	})

	Describe("StringCmd", func() {
		It("is implemented by iredis.StringCmdWrap", func() {
			var _ iredis.StringCmd = new(iredis.StringCmdWrap)
		})

		It("is implemented by iredis.StringCmdFake", func() {
			var _ iredis.StringCmd = new(iredis.StringCmdFake)
		})
	})
})
