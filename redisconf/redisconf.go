package redisconf

import "fmt"

//Directive describes the methods for `redis.conf` directives
type Directive interface {
	String() string
	IsValid() error
}

//RedisConf is a representation of `redis.conf`
type RedisConf []Directive

//Config is a config value for a `redis.conf`
type Config struct {
	Name  string
	Value string
}

//NewConfig is a conventient way to initialise a Config
func NewConfig(name, value string) Config {
	return Config{Name: name, Value: value}
}

func (config Config) String() string {
	return fmt.Sprintf("%s %s", config.Name, config.Value)
}

//IsValid returns an error if it considered invalid by Redis
func (config Config) IsValid() error {
	return nil
}

//RenameCommand is a command alias for a Redis command
type RenameCommand struct {
	Command string
	Alias   string
}

//NewRenameCommand is a conventient way to initialise a RenameCommand
func NewRenameCommand(command, alias string) RenameCommand {
	return RenameCommand{Command: command, Alias: alias}
}

func (rename RenameCommand) String() string {
	return fmt.Sprintf("rename-command %s %s", rename.Command, rename.aliasString())
}

//IsValid returns an error if it considered invalid by Redis
func (rename RenameCommand) IsValid() error {
	return nil
}

func (rename RenameCommand) aliasString() string {
	if rename.Alias == "" {
		return `""`
	}
	return rename.Alias
}
