package tunnel

import (
	"github.com/garyburd/redigo/redis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/ssh"
)

var _ = Describe("tunnel", func() {
	It("pings a redis instance behind a bastion", func() {
		// SPIKE AHEAD
		// This is a spike to test pinging a Redis instance behind a bastion
		// local machine --- tcp ---> local server --- ssh ---> bastion host --- tcp ---> redis
		// Create a bridge network to hold these containers: `docker network create --driver bridge test-go-forth`
		// In order to run the baston host: `docker run -i --network=test-go-forth -p 8001:22/tcp --user root -t cflondonservices/redisutils`
		// Start the ssh server on the bastion: `/etc/init.d/ssh start`
		// In order to run the redis host: `docker run -i --network=test-go-forth  --expose 6379 -t cflondonservices/redisutils`
		// Run this test

		// TODO
		// Make this work in concourse (docker inside docker so you can docker your docker)

		local := Endpoint{Host: "localhost", Port: 8005}
		server := Endpoint{Host: sshHost, Port: sshPort}
		remote := Endpoint{Host: redisHost, Port: redisPort}

		tunnel := &SSHTunnel{
			Local:  local,
			Server: server,
			Remote: remote,
			Config: &ssh.ClientConfig{User: "vcap", Auth: []ssh.AuthMethod{ssh.Password("funky92horse")}},
		}

		go tunnel.Start()
		conn, err := redis.Dial("tcp", "localhost:8005")
		Expect(err).NotTo(HaveOccurred())
		defer conn.Close()
		reply, err := conn.Do("ping")
		Expect(err).NotTo(HaveOccurred())
		Expect(reply).To(Equal("PONG"))
	})
})
