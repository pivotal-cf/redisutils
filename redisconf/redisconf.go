package redisconf

import "fmt"

//Directive describes the methods for `redis.conf` directives
type Directive interface {
	String() string
	Validate() error
}

//RedisConf is a representation of `redis.conf`
type RedisConf []Directive

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

//Validate returns an error if it considered invalid by Redis
func (rename RenameCommand) Validate() error {
	return nil
}

func (rename RenameCommand) aliasString() string {
	if rename.Alias == "" {
		return `""`
	}
	return rename.Alias
}
