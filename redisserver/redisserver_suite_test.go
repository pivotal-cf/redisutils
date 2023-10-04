package redisserver

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRedisServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "redisserver suite")
}
