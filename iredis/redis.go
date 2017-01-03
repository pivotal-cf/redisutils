package iredis

import redis "gopkg.in/redis.v5"

type IRedis interface {
	NewClient(*redis.Options) IClient
}

//RedisWrap is a wrapper around redis that implements iredis.IRedis
type RedisWrap struct{}

//NewClient is a wrapper around redis.NewClient()
func (*RedisWrap) NewClient(opt *redis.Options) IClient {
	return &ClientWrap{client: redis.NewClient(opt)}
}
