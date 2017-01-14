package redisconf

import (
	"fmt"
	"os"

	"github.com/BooleanCat/igo/iioutil"
)

//RedisConf is a representation of a `redis.conf` file
type RedisConf interface {
	Get(string) string
	Set(string, string)
	GetRenameCommand(string) string
	SetRenameCommand(string, string)
	Save(string) error
}

//Conf is the implementation of RedisConf
type Conf struct {
	configs map[string]string

	ioutil iioutil.Ioutil
}

//New is the correct way to initialise a RedisConf
func New() *Conf {
	return &Conf{
		configs: map[string]string{
			"host": "localhost",
			"port": "6379",
		},
		ioutil: iioutil.New(),
	}
}

//Get a redis config value
func (c *Conf) Get(config string) string {
	return c.configs[config]
}

//Set a redis config value
func (c *Conf) Set(config, value string) {
	c.configs[config] = value
}

//GetRenameCommand a redis config value
func (c *Conf) GetRenameCommand(string) string {
	return ""
}

//SetRenameCommand a redis config value
func (c *Conf) SetRenameCommand(string, string) {}

//Save the config to disk
func (c *Conf) Save(path string) error {
	contents := c.encode()
	return c.ioutil.WriteFile(path, []byte(contents), os.ModePerm)
}

func (c *Conf) encode() (contents string) {
	for config, value := range c.configs {
		line := fmt.Sprintf("%s %s", config, value)
		contents = contents + line + "\n"
	}
	return
}
