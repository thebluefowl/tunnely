package store

import "github.com/thebluefowl/tunnely/tunnel"

// DomainMap stores the subdomain and the corresponding connection identifier.
type TunnelStore interface {
	GetTunnelByDomain(string) *tunnel.Tunnel
	GetTunnelByID(string) *tunnel.Tunnel
	GetDomains() []string
	GetIDs() []string
	AddTunnel(*tunnel.Tunnel)
}
