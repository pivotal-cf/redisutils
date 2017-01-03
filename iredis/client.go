package iredis

import redis "gopkg.in/redis.v5"

type IClient interface {
	Close() error
	Ping() IStatusCmd
}

//ClientWrap is a wrapper around redis that implements iredis.IClient
type ClientWrap struct {
	client *redis.Client
}

//Close is a wrapper around redis.Client.Close()
func (clientWrap *ClientWrap) Close() error {
	return clientWrap.client.Close()
}

//Ping is a wrapper around redis.Client.Ping()
func (clientWrap *ClientWrap) Ping() IStatusCmd {
	return &StatusCmdWrap{statusCmd: clientWrap.client.Ping()}
}
