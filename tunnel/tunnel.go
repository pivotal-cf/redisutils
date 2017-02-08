package tunnel

import (
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/BooleanCat/igo/inet"
	"golang.org/x/crypto/ssh"
)

//Endpoint represents a network endpoint
type Endpoint struct {
	Host string
	Port int
}

func (endpoint Endpoint) String() string {
	return fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
}

//SSHTunnel is used to tunnel through a Server into a Remote host
type SSHTunnel struct {
	Local  Endpoint
	Server Endpoint
	Remote Endpoint
	Config *ssh.ClientConfig

	err     error
	errLock *sync.Mutex
	net     inet.Net
}

//New is the correct way to initialise a SSHTunnel
func New(local, server, remote Endpoint, config *ssh.ClientConfig) *SSHTunnel {
	return &SSHTunnel{
		Local:  local,
		Server: server,
		Remote: remote,
		Config: config,

		errLock: new(sync.Mutex),
		net:     inet.New(),
	}
}

//Start tunneling through Server to Remote
func (tunnel *SSHTunnel) Start() {
	listener, err := tunnel.net.Listen("tcp", tunnel.Local.String())
	if err != nil {
		tunnel.setErr(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			tunnel.setErr(err)
			return
		}
		go tunnel.forward(conn)
	}
}

//GetErr gets the error if one has occurred within Start
func (tunnel *SSHTunnel) GetErr() error {
	tunnel.errLock.Lock()
	defer tunnel.errLock.Unlock()
	return tunnel.err
}

func (tunnel *SSHTunnel) setErr(err error) {
	tunnel.errLock.Lock()
	defer tunnel.errLock.Unlock()
	tunnel.err = err
}

func (tunnel *SSHTunnel) forward(localConn net.Conn) {
	serverConn, err := ssh.Dial("tcp", tunnel.Server.String(), tunnel.Config)
	if err != nil {
		fmt.Printf("Server dial error: %s\n", err)
		return
	}

	remoteConn, err := serverConn.Dial("tcp", tunnel.Remote.String())
	if err != nil {
		fmt.Printf("Remote dial error: %s\n", err)
		return
	}

	copyConn := func(writer, reader net.Conn) {
		_, err := io.Copy(writer, reader)
		if err != nil {
			fmt.Printf("io.Copy error: %s", err)
		}
	}

	go copyConn(localConn, remoteConn)
	go copyConn(remoteConn, localConn)
}
