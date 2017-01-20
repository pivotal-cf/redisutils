package iredis

import redis "gopkg.in/redis.v5"

type scripter interface {
}

//Script is an interface around redis.Script
type Script interface {
	Eval(scripter, []string, ...interface{}) *redis.Cmd
	EvalSha(scripter, []string, ...interface{}) *redis.Cmd
	Exists(scripter) *BoolSliceCmd
	Load(scripter) *StringCmd
	Run(scripter, []string, ...interface{}) *redis.Cmd
}
