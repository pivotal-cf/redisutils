package tunnel

import (
	"errors"

	"github.com/BooleanCat/igo/inet"
	"github.com/garyburd/redigo/redis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("tunnel", func() {
	var netFake *inet.Fake

	BeforeEach(func() {
		netFake = inet.NewFake()
	})

	Describe("Endpoint", func() {
		Describe("#String", func() {
			It("is represented as `host:port`", func() {
				endpoint := Endpoint{Host: "localhost", Port: 80}
				Expect(endpoint.String()).To(Equal("localhost:80"))
			})
		})
	})

	Describe("#SSHTunnel", func() {
		var (
			sshTunnel *SSHTunnel
			local     = Endpoint{Host: "localhost", Port: 8005}
		)

		BeforeEach(func() {
			server := Endpoint{Host: sshHost, Port: sshPort}
			remote := Endpoint{Host: redisHost, Port: redisPort}
			sshTunnel = makeSSHTunnel(local, server, remote)
		})

		Describe("#Start", func() {
			JustBeforeEach(func() {
				go sshTunnel.Start()
			})

			It("creates a TCP tunnel", func() {
				connection, err := redis.Dial("tcp", local.String())
				Expect(err).NotTo(HaveOccurred())
				defer connection.Close()

				By("being pingable")
				reply, err := connection.Do("ping")
				Expect(err).NotTo(HaveOccurred())
				Expect(reply).To(Equal("PONG"))

				By("not returning an error")
				Expect(sshTunnel.GetErr()).NotTo(HaveOccurred())
			})

			Context("when creating a local server fails", func() {
				listenErr := errors.New("listen failed")

				BeforeEach(func() {
					netFake.ListenReturns(nil, listenErr)
					sshTunnel.net = netFake
				})

				It("sets an error", func() {
					getErr := func() error { return sshTunnel.GetErr() }
					Eventually(getErr).Should(MatchError(listenErr))
				})
			})
		})
	})
})
