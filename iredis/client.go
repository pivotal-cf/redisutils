package iredis

import redis "gopkg.in/redis.v5"

//Client is an interface around redis.Client
type Client interface {
	Close() error
	Ping() StatusCmd
	BgRewriteAOF() StatusCmd
	Info(...string) StringCmd
}

//ClientWrap is a wrapper around redis that implements iredis.Client
type ClientWrap struct {
	client *redis.Client
}

//Close is a wrapper around redis.Client.Close()
func (clientWrap *ClientWrap) Close() error {
	return clientWrap.client.Close()
}

//Ping is a wrapper around redis.Client.Ping()
func (clientWrap *ClientWrap) Ping() StatusCmd {
	return &StatusCmdWrap{statusCmd: clientWrap.client.Ping()}
}

//BgRewriteAOF is a wrapper around redis.Client.BgRewriteAOF()
func (clientWrap *ClientWrap) BgRewriteAOF() StatusCmd {
	return &StatusCmdWrap{statusCmd: clientWrap.client.BgRewriteAOF()}
}

//Info is a wrapper around redis.Client.Info()
func (clientWrap *ClientWrap) Info(section ...string) StringCmd {
	return &StringCmdWrap{stringCmd: clientWrap.client.Info(section...)}
}
