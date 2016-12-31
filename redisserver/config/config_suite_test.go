package config

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

func createTempDir() string {
	tempDir, err := ioutil.TempDir("", "")
	Expect(err).NotTo(HaveOccurred())
	return tempDir
}

func removeAllIfTemp(dir string) {
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
