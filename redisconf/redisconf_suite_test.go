package redisconf_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pivotal-cf/redisutils/redisconf"
)

func TestRedisconf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "redisconf suite")
}

func newRedisConf(directives ...redisconf.Directive) redisconf.RedisConf {
	redisConf, err := redisconf.New(directives...)
	Expect(err).NotTo(HaveOccurred())
	return redisConf
}

func newDirective(keyword string, args ...string) redisconf.Directive {
	directive, err := redisconf.NewDirective(keyword, args...)
	Expect(err).NotTo(HaveOccurred())
	return directive
}
