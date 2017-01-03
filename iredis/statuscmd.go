package iredis

import redis "gopkg.in/redis.v5"

//StatusCmd is an interface around redis.StatusCmd
type StatusCmd interface {
	Result() (string, error)
}

//StatusCmdWrap is a wrapper around redis that implements iredis.StatusCmd
type StatusCmdWrap struct {
	statusCmd *redis.StatusCmd
}

//Result is a wrapper around redis.StatusCmd.Result()
func (scw *StatusCmdWrap) Result() (string, error) {
	return scw.statusCmd.Result()
}
