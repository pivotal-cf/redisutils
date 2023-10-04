package redis_test

import (
	redigo "github.com/gomodule/redigo/redis"
	. "github.com/onsi/ginkgo/v2"

	"github.com/pivotal-cf/redisutils/redis"
)

var _ = Describe("interfaces", func() {
	Describe("Redis", func() {
		It("is implemented by redis.Real", func() {
			var _ redis.Redis = redis.New()
		})

		It("is implemented by redis.Fake", func() {
			var _ redis.Redis = new(redis.Fake)
		})
	})

	Describe("Conn", func() {
		It("is implemented by redis.ConnFake", func() {
			var _ redigo.Conn = new(redis.ConnFake)
		})
	})
})
