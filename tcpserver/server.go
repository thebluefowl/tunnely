package tcpserver

import (
	"bufio"
	"fmt"
	"net"

	"github.com/thebluefowl/tunnely/store"
	"github.com/thebluefowl/tunnely/tunnel"
)

const (
	CmdInit  = "INI"
	CmdClose = "FIN"
)

type TCPServer struct {
	Port  int64
	store store.DomainStore
}

func NewTCPServer(port int64) *TCPServer {
	server := &TCPServer{
		Port: port,
	}
	server.store = store.GetInMemoryDomainStore()
	return server
}

func (ts *TCPServer) Listen() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", ts.Port))
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}

		go func() {
			for {
				reader := bufio.NewReader(c)
				bytes, err := reader.Peek(3)
				switch string(bytes) {
				case CmdInit:
					reader.Discard(3)
					bytes, err = reader.Peek(20)
					id := string(bytes)
					reader.Discard(20)
					ts.store.AddTunnel(tunnel.NewTunnel(id, c))
				}
				x, _, _ := reader.ReadLine()

				if err != nil {

					return
				}
				fmt.Println(string(x))
			}
		}()
	}
}
