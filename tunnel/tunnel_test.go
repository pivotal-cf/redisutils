package tunnel

import (
	"github.com/garyburd/redigo/redis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("tunnel", func() {
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
			var connection redis.Conn

			BeforeEach(func() {
				go sshTunnel.Start()

				var dialErr error
				connection, dialErr = redis.Dial("tcp", local.String())
				Expect(dialErr).NotTo(HaveOccurred())
			})

			AfterEach(func() {
				err := connection.Close()
				Expect(err).NotTo(HaveOccurred())
			})

			It("creates a TCP tunnel", func() {
				reply, err := connection.Do("ping")
				Expect(err).NotTo(HaveOccurred())
				Expect(reply).To(Equal("PONG"))
			})

			It("does not set an error", func() {
				Expect(sshTunnel.GetErr()).NotTo(HaveOccurred())
			})
		})
	})
})
