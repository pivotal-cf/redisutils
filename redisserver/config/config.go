package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/BooleanCat/igo/iioutil"
)

//Config is a go representation of `redis.conf`
type Config []fmt.Stringer

func (config Config) String() string {
	return strings.Join(config.getStrings(), "\n")
}

//ToFile writes the stringyfied config to a `redis.conf` file at path
func (config Config) ToFile(path string) error {
	return config.toFile(path, iioutil.New())
}

func (config Config) toFile(path string, ioutil iioutil.Ioutil) error {
	return ioutil.WriteFile(path, []byte(config.String()), os.ModePerm)
}

func (config Config) getStrings() []string {
	confStrings := []string{}
	for _, conf := range config {
		confStrings = append(confStrings, conf.String())
	}
	return confStrings
}

//Simple representents configs like `save 900 1`
type Simple struct {
	Key   string
	Value string
}

//NewSimple is a shorthand way to create a Simple
func NewSimple(key, value string) Simple {
	return Simple{Key: key, Value: value}
}

func (config Simple) String() string {
	return fmt.Sprintf("%s %s", config.Key, config.Value)
}

//RenameCommand represents configs like `rename-command CONFIG foo`
type RenameCommand struct {
	Command string
	Alias   string
}

//NewRenameCommand is a shorthand way to create a RenameCommand
func NewRenameCommand(command, alias string) RenameCommand {
	return RenameCommand{Command: command, Alias: alias}
}

func (config RenameCommand) String() string {
	alias := `""`
	if config.Alias != "" {
		alias = config.Alias
	}

	return fmt.Sprintf("rename-command %s %s", config.Command, alias)
}
