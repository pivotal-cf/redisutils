package redisserver_integration_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRedisServerIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RedisServer Integration Suite")
}
