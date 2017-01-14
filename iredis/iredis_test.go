package iredis_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/pivotal-cf/redisutils/iredis"
)

var _ = Describe("interfaces", func() {
	Describe("Redis", func() {
		It("is implemented by iredis.RedisWrap", func() {
			redisWrap := new(iredis.RedisWrap)
			var _ = iredis.Redis(redisWrap)
		})

		It("is implemented by iredis.RedisFake", func() {
			redisFake := new(iredis.RedisFake)
			var _ = iredis.Redis(redisFake)
		})
	})

	Describe("Client", func() {
		It("is implemented by iredis.ClientWrap", func() {
			clientWrap := new(iredis.ClientWrap)
			var _ = iredis.Client(clientWrap)
		})

		It("is implemented by iredis.ClientFake", func() {
			clientFake := new(iredis.ClientFake)
			var _ = iredis.Client(clientFake)
		})
	})

	Describe("StatusCmd", func() {
		It("is implemented by iredis.StatusCmdWrap", func() {
			statusCmdWrap := new(iredis.StatusCmdWrap)
			var _ = iredis.StatusCmd(statusCmdWrap)
		})

		It("is implemented by iredis.StatusCmdFake", func() {
			statusCmdFake := new(iredis.StatusCmdFake)
			var _ = iredis.StatusCmd(statusCmdFake)
		})
	})

	Describe("StringCmd", func() {
		It("is implemented by iredis.StringCmdWrap", func() {
			stringCmdWrap := new(iredis.StringCmdWrap)
			var _ = iredis.StringCmd(stringCmdWrap)
		})

		It("is implemented by iredis.StringCmdFake", func() {
			stringCmdFake := new(iredis.StringCmdFake)
			var _ = iredis.StringCmd(stringCmdFake)
		})
	})
})
