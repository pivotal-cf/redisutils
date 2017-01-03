package iredis

import redis "gopkg.in/redis.v5"

type Redis interface {
	NewClient(*redis.Options) Client
}

//RedisWrap is a wrapper around redis that implements iredis.Redis
type RedisWrap struct{}

//NewClient is a wrapper around redis.NewClient()
func (*RedisWrap) NewClient(opt *redis.Options) Client {
	return &ClientWrap{client: redis.NewClient(opt)}
}
