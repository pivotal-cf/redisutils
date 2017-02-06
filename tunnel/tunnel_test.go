package tunnel

import (
	"golang.org/x/crypto/ssh"

	"github.com/garyburd/redigo/redis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
		// Have the docker image always start the ssh-server on start up
		// Have the test suite create and destroy the network
		// Have the test suite launch both docker containers
		// Have the test suite kill both docker containers
		// Make this work in concourse (docker inside docker so you can docker your docker)

		local := Endpoint{Host: "localhost", Port: 8005}
		server := Endpoint{Host: "localhost", Port: 8001}
		remote := Endpoint{Host: "172.18.0.3", Port: 6379}

		tunnel := &SSHTunnel{
			Local:  local,
			Server: server,
			Remote: remote,
			Config: &ssh.ClientConfig{User: "vcap", Auth: []ssh.AuthMethod{ssh.Password("vcap")}},
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
