package redisserver

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRedisServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RedisServer Suite")
}
