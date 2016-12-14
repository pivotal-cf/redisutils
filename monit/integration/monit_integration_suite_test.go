package monit_integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMonitIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Monit Integration Suite")
}
