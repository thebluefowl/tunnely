package tunnel

import (
	"crypto/rand"
	"net"
)

type Tunnel struct {
	ID        string
	Subdomain string
	Conn      net.Conn
}

func NewTunnel(id string, conn net.Conn) *Tunnel {
	return &Tunnel{
		ID:        id,
		Subdomain: randString(8),
		Conn:      conn,
	}
}

func randString(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return string(b)
}
