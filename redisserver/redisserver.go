package redisserver

//RedisServer is a controller for `redis-server`
type RedisServer interface {
	Start() error
	Stop() error
}

//Server is a controller for `redis-server`
type Server struct{}

//New is the correct way to initialise a Server
func New() *Server {
	return new(Server)
}

//Start launches a `redis-server`
func (server *Server) Start() error {
	return nil
}

//Stop kills a running `redis-server`
func (server *Server) Stop() error {
	return nil
}
