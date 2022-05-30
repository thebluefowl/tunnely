package tcpserver

import (
	"bufio"
	"fmt"
	"net"

	"github.com/thebluefowl/tunnely/store"
	"github.com/thebluefowl/tunnely/tunnel"
)

const (
	CmdInit        = "INIT"
	CmdClose       = "FINI"
	CmdListDomains = "LSTD"
	CmdPing        = "PING"
)

type TCPServer struct {
	Port  int64
	store store.TunnelStore
}

func NewTCPServer(port int64, store store.TunnelStore) *TCPServer {
	server := &TCPServer{
		Port: port,
	}
	server.store = store
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
				bytes, err := reader.Peek(4)
				if err != nil {
					fmt.Println(err)
					continue
				}
				switch string(bytes) {
				case CmdInit:
					reader.Discard(4)
					bytes, err = reader.Peek(20)
					if err != nil {
						fmt.Println(err)
						continue
					}
					id := string(bytes)
					reader.Discard(20)
					ts.store.AddTunnel(tunnel.NewTunnel(id, c))
					return
				case CmdListDomains:
					fmt.Println(ts.store.GetIDs())
				case CmdPing:
					reader.Discard(4)
					bytes, err := reader.Peek(20)
					if err != nil {
						fmt.Println(err)
						continue
					}
					id := string(bytes)
					tunnel := ts.store.GetTunnelByID(id)
					_, err = tunnel.Conn.Write([]byte("Pong"))
					if err != nil {
						fmt.Println(err)
						return
					}
				}

				x, _, _ := reader.ReadLine()
				fmt.Println(string(x))
			}
		}()
		fmt.Println("here")
	}
}
