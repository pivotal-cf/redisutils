package iioutil

import (
	"io"
	"io/ioutil"
	"os"
)

//Ioutil is an interface around ioutil
type Ioutil interface {
	ReadAll(io.Reader) ([]byte, error)
	ReadDir(string) ([]os.FileInfo, error)
	ReadFile(string) ([]byte, error)
	TempDir(string, string) (string, error)
	TempFile(string, string) (*os.File, error)
	WriteFile(string, []byte, os.FileMode) error
}

//IoutilWrap is a wrapper around os that implements iioutil.Ioutil
type IoutilWrap struct{}

//ReadAll is a wrapper around ioutil.ReadAll()
func (iow *IoutilWrap) ReadAll(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}

//ReadDir is a wrapper around ioutil.ReadDir()
func (iow *IoutilWrap) ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

//ReadFile is a wrapper around ioutil.ReadFile()
func (iow *IoutilWrap) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

//TempDir is a wrapper around ioutil.TempDir()
func (iow *IoutilWrap) TempDir(dir, prefix string) (string, error) {
	return ioutil.TempDir(dir, prefix)
}

//TempFile is a wrapper around ioutil.TempDir()
func (iow *IoutilWrap) TempFile(dir, prefix string) (*os.File, error) {
	return ioutil.TempFile(dir, prefix)
}

//WriteFile is a wrapper around ioutil.TempDir()
func (iow *IoutilWrap) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}
