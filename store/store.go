package store

import "github.com/thebluefowl/tunnely/tunnel"

// DomainMap stores the subdomain and the corresponding connection identifier.
type DomainStore interface {
	GetTunnel(string) *tunnel.Tunnel
	AddTunnel(*tunnel.Tunnel)
}
