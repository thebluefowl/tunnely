package tunnel

import (
	"math/rand"
	"net"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

type Tunnel struct {
	ID        string
	Subdomain string
	Conn      net.Conn
}

func NewTunnel(id string, conn net.Conn) *Tunnel {
	return &Tunnel{
		ID: id,
		// Subdomain: randomString(7),
		Subdomain: "abc",
		Conn:      conn,
	}
}

// randomString generates a random alphabetical string of length `n`
func randomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn((len(charset)))]
	}
	return string(b)
}
