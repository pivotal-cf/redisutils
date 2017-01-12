package config

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/BooleanCat/igo/iioutil"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("config", func() {
	var (
		config      fmt.Stringer
		stringyfied string
		ioutilFake  *iioutil.Fake
	)

	BeforeEach(func() {
		ioutilFake = iioutil.NewFake()
	})

	JustBeforeEach(func() {
		stringyfied = fmt.Sprintf("%s", config)
	})

	Describe("Simple", func() {
		BeforeEach(func() {
			config = NewSimple("foo", "bar")
		})

		It("stringyfies correctly", func() {
			Expect(stringyfied).To(Equal("foo bar"))
		})
	})

	Describe("RenameCommand", func() {
		BeforeEach(func() {
			config = NewRenameCommand("bar", "baz")
		})

		It("stringyfies correctly", func() {
			Expect(stringyfied).To(Equal("rename-command bar baz"))
		})

		Context("when alias is empty", func() {
			BeforeEach(func() {
				config = NewRenameCommand("bar", "")
			})

			It("represents the alias as double quotes", func() {
				Expect(stringyfied).To(Equal(`rename-command bar ""`))
			})
		})
	})

	Describe("Config", func() {
		var testConfig Config

		BeforeEach(func() {
			testConfig = Config{
				NewSimple("foo", "bar"),
				NewRenameCommand("baz", ""),
				NewSimple("mip", "map"),
				NewRenameCommand("bit", "bot"),
			}
			config = testConfig
		})

		It("stringyfies correctly", func() {
			expected := strings.Join([]string{
				"foo bar",
				`rename-command baz ""`,
				"mip map",
				"rename-command bit bot",
			}, "\n")

			Expect(stringyfied).To(Equal(expected))
		})

		Describe("ToFile", func() {
			var (
				toFileErr  error
				tempDir    string
				configPath string
			)

			BeforeEach(func() {
				tempDir = createTempDir()
				configPath = filepath.Join(tempDir, "redis.conf")
				toFileErr = testConfig.ToFile(configPath)
			})

			AfterEach(func() {
				removeAllIfTemp(tempDir)
			})

			It("does not return an error", func() {
				Expect(toFileErr).NotTo(HaveOccurred())
			})

			It("creates a redis.conf file", func() {
				Expect(configPath).To(BeAnExistingFile())
			})

			It("writes the correct config to the file", func() {
				conf := readFile(configPath)
				Expect(conf).To(Equal(config.String()))
			})

			Context("when WriteFile returns an error", func() {
				writeFileErr := errors.New("WriteFile failed")

				BeforeEach(func() {
					ioutilFake.WriteFileReturns(writeFileErr)
					toFileErr = testConfig.toFile(configPath, ioutilFake)
				})

				It("returns the error", func() {
					Expect(toFileErr).To(MatchError(writeFileErr))
				})
			})
		})
	})
})
