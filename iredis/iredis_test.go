package iredis_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/pivotal-cf/redisutils/iredis"
)

var _ = Describe("interfaces", func() {
	Describe("Redis", func() {
		It("is implemented by iredis.Real", func() {
			var _ iredis.Redis = iredis.New()
		})

		It("is implemented by iredis.Fake", func() {
			var _ iredis.Redis = iredis.NewFake()
		})
	})

	Describe("Client", func() {
		It("is implemented by iredis.ClientReal", func() {
			var _ iredis.Client = new(iredis.ClientReal)
		})

		It("is implemented by iredis.ClientFake", func() {
			var _ iredis.Client = iredis.NewClientFake()
		})
	})

	Describe("StatusCmd", func() {
		It("is implemented by iredis.StatusCmdReal", func() {
			var _ iredis.StatusCmd = new(iredis.StatusCmdReal)
		})

		It("is implemented by iredis.StatusCmdFake", func() {
			var _ iredis.StatusCmd = iredis.NewStatusCmdFake()
		})
	})

	Describe("StringCmd", func() {
		It("is implemented by iredis.StringCmdReal", func() {
			var _ iredis.StringCmd = new(iredis.StringCmdReal)
		})

		It("is implemented by iredis.StringCmdFake", func() {
			var _ iredis.StringCmd = iredis.NewStringCmdFake()
		})
	})

	Describe("BoolSliceCmd", func() {
		It("is implemented by iredis.BoolSliceCmdReal", func() {
			var _ iredis.BoolSliceCmd = new(iredis.BoolSliceCmdReal)
		})

		It("is implemented by iredis.BoolSliceCmdFake", func() {
			var _ iredis.BoolSliceCmd = iredis.NewBoolSliceCmdFake()
		})
	})

	Describe("Script", func() {
		It("is implemented by iredis.ScriptReal", func() {
			var _ iredis.Script = new(iredis.ScriptReal)
		})

		It("is implemented by iredis.ScriptFake", func() {
			var _ iredis.Script = iredis.NewScriptFake()
		})
	})
})
