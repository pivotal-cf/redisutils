package iredis

import redis "gopkg.in/redis.v5"

type StringCmd interface {
	Result() (string, error)
}

//StringCmdWrap is a wrapper around redis that implements iredis.StringCmd
type StringCmdWrap struct {
	stringCmd *redis.StringCmd
}

//Result is a wrapper around redis.StringCmd.Result()
func (scw *StringCmdWrap) Result() (string, error) {
	return scw.stringCmd.Result()
}
