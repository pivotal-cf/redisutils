package iredis

import . "github.com/onsi/ginkgo"

var _ = Describe("interfaces", func() {
	Describe("Redis", func() {
		It("is implemented by iredis.RedisWrap", func() {
			redisWrap := new(RedisWrap)
			var _ = Redis(redisWrap)
		})

		It("is implemented by iredis.RedisFake", func() {
			redisFake := new(RedisFake)
			var _ = Redis(redisFake)
		})
	})

	Describe("Client", func() {
		It("is implemented by iredis.ClientWrap", func() {
			clientWrap := new(ClientWrap)
			var _ = Client(clientWrap)
		})

		It("is implemented by iredis.ClientFake", func() {
			clientFake := new(ClientFake)
			var _ = Client(clientFake)
		})
	})

	Describe("StatusCmd", func() {
		It("is implemented by iredis.StatusCmdWrap", func() {
			statusCmdWrap := new(StatusCmdWrap)
			var _ = StatusCmd(statusCmdWrap)
		})

		It("is implemented by iredis.StatusCmdFake", func() {
			statusCmdFake := new(StatusCmdFake)
			var _ = StatusCmd(statusCmdFake)
		})
	})

	Describe("StringCmd", func() {
		It("is implemented by iredis.StringCmdWrap", func() {
			stringCmdWrap := new(StringCmdWrap)
			var _ = StringCmd(stringCmdWrap)
		})

		It("is implemented by iredis.StringCmdFake", func() {
			stringCmdFake := new(StringCmdFake)
			var _ = StringCmd(stringCmdFake)
		})
	})
})
