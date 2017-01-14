package iredis

import redis "gopkg.in/redis.v5"

//Client is an interface around redis.Client
type Client interface {
	Close() error
	Ping() StatusCmd
	BgRewriteAOF() StatusCmd
	Info(...string) StringCmd
}

//ClientReal is a wrapper around redis that implements iredis.Client
type ClientReal struct {
	client *redis.Client
}

//NewClient is a wrapper around redis.NewClient()
func (*Real) NewClient(opt *redis.Options) Client {
	return &ClientReal{client: redis.NewClient(opt)}
}

//Close is a wrapper around redis.Client.Close()
func (clientWrap *ClientReal) Close() error {
	return clientWrap.client.Close()
}

//Ping is a wrapper around redis.Client.Ping()
func (clientWrap *ClientReal) Ping() StatusCmd {
	return &StatusCmdReal{statusCmd: clientWrap.client.Ping()}
}

//BgRewriteAOF is a wrapper around redis.Client.BgRewriteAOF()
func (clientWrap *ClientReal) BgRewriteAOF() StatusCmd {
	return &StatusCmdReal{statusCmd: clientWrap.client.BgRewriteAOF()}
}

//Info is a wrapper around redis.Client.Info()
func (clientWrap *ClientReal) Info(section ...string) StringCmd {
	return &StringCmdReal{stringCmd: clientWrap.client.Info(section...)}
}
