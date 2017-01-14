package iredis

import redis "gopkg.in/redis.v5"

//StringCmd is an interface around redis.StringCmd
type StringCmd interface {
	Result() (string, error)
}

//StringCmdReal is a wrapper around redis that implements iredis.StringCmd
type StringCmdReal struct {
	stringCmd *redis.StringCmd
}

//Result is a wrapper around redis.StringCmd.Result()
func (cmd *StringCmdReal) Result() (string, error) {
	return cmd.stringCmd.Result()
}
