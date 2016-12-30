package redisserver

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("redisserver", func() {
	var server *Server

	BeforeEach(func() {
		server = New()
	})

	Describe("#Start", func() {
		var startErr error

		BeforeEach(func() {
			startErr = server.Start()
		})

		It("does not return an error", func() {
			Expect(startErr).NotTo(HaveOccurred())
		})
	})

	Describe("#Stop", func() {
		var stopErr error

		BeforeEach(func() {
			stopErr = server.Stop()
		})

		It("does not return an error", func() {
			Expect(stopErr).NotTo(HaveOccurred())
		})
	})
})
