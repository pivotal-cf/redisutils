package iredis_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIRedis(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IRedis Suite")
}
