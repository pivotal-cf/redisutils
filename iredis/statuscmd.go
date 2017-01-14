package iredis

import redis "gopkg.in/redis.v5"

//StatusCmd is an interface around redis.StatusCmd
type StatusCmd interface {
	Result() (string, error)
}

//StatusCmdReal is a wrapper around redis that implements iredis.StatusCmd
type StatusCmdReal struct {
	statusCmd *redis.StatusCmd
}

//NewStatusCmd is a wrapper around redis.NewStatusCmd()
func NewStatusCmd(args ...interface{}) StatusCmd {
	return &StatusCmdReal{statusCmd: redis.NewStatusCmd(args...)}
}

//Result is a wrapper around redis.StatusCmd.Result()
func (cmd *StatusCmdReal) Result() (string, error) {
	return cmd.statusCmd.Result()
}
