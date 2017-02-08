package inet

import "net"

//Net is an interface around net
type Net interface {
	Listen(protocol, laddr string) (net.Listener, error)
}

//Real is a wrapper around net that implements inet.Net
type Real struct{}

//New creates a struct that behaves like the os package
func New() *Real {
	return new(Real)
}

//Listen is a wrapper around net.Listen()
func (*Real) Listen(protocol, laddr string) (net.Listener, error) {
	return net.Listen(protocol, laddr)
}
