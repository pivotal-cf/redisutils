package redisconf

import (
	"errors"
	"path/filepath"

	"github.com/BooleanCat/igo/iioutil"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("redisconf", func() {
	var (
		redisConf  *Conf
		ioutilFake *iioutil.Fake
	)

	BeforeEach(func() {
		ioutilFake = iioutil.NewFake()
		redisConf = New()
	})

	Describe("#Get", func() {
		It("defaults host correctly", func() {
			Expect(redisConf.Get("host")).To(Equal("localhost"))
		})

		It("defaults port correctly", func() {
			Expect(redisConf.Get("port")).To(Equal("6379"))
		})

		It("defaults requirepass correctly", func() {
			Expect(redisConf.Get("requirepass")).To(Equal(""))
		})
	})

	Describe("#Set", func() {
		It("sets a value", func() {
			redisConf.Set("foo", "bar")
			Expect(redisConf.Get("foo")).To(Equal("bar"))
		})

		It("updates existing configs", func() {
			redisConf.Set("host", "example.com")
			Expect(redisConf.Get("host")).To(Equal("example.com"))
		})
	})

	Describe("#Save", func() {
		var (
			saveErr       error
			tempDir       string
			redisConfPath string
		)

		BeforeEach(func() {
			tempDir = createTempDir()
			redisConfPath = filepath.Join(tempDir, "redis.conf")
		})

		JustBeforeEach(func() {
			saveErr = redisConf.Save(redisConfPath)
		})

		AfterEach(func() {
			removeTempDir(tempDir)
		})

		It("does not return an error", func() {
			Expect(saveErr).NotTo(HaveOccurred())
		})

		It("creates a file on disk", func() {
			Expect(redisConfPath).To(BeAnExistingFile())
		})

		Context("when ioutil.WriteFile returns an error", func() {
			writeFileErr := errors.New("WriteFile failed")

			BeforeEach(func() {
				redisConf.ioutil = ioutilFake
				ioutilFake.WriteFileReturns(writeFileErr)
			})

			It("returns the error", func() {
				Expect(saveErr).To(MatchError(writeFileErr))
			})
		})
	})
})
