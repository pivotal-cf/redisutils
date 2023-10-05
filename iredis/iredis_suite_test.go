package iredis_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIRedis(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "iredis suite")
}
