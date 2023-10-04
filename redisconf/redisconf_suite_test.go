package redisconf

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRedisconf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "redisconf suite")
}
