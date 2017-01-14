package redisconf

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
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

func readFile(path string) string {
	contents, err := ioutil.ReadFile(path)
	Expect(err).NotTo(HaveOccurred())
	return string(contents)
}

func containLine(line string) types.GomegaMatcher {
	return MatchRegexp(fmt.Sprintf("(?m)^%s$", line))
}
