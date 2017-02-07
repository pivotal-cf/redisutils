package tunnel

import (
	"os"
	"strconv"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	redisHost string
	redisPort int
	sshHost   string
	sshPort   int
)

func TestTunnel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "tunnel suite")
}

var _ = BeforeSuite(func() {
	redisHost = mustGetenv("REDIS_HOST")
	redisPort = mustGetenvInt("REDIS_PORT")
	sshHost = mustGetenv("SSH_HOST")
	sshPort = mustGetenvInt("SSH_PORT")
})

func mustGetenv(env string) string {
	value := os.Getenv(env)
	Expect(value).NotTo(Equal(""))
	return value
}

func mustGetenvInt(env string) int {
	value := mustGetenv(env)
	intValue, err := strconv.Atoi(value)
	Expect(err).NotTo(HaveOccurred())
	return intValue
}
