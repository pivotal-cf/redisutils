package redisconf

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRedisconf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "redisconf suite")
}

func createTempDir() string {
	dir, err := ioutil.TempDir("", "")
	Expect(err).NotTo(HaveOccurred())
	return dir
}

func removeTempDir(dir string) {
	if strings.HasPrefix(dir, os.TempDir()) {
		os.RemoveAll(dir)
		Expect(dir).NotTo(BeAnExistingFile())
	}
}
