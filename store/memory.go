package store

import (
	"fmt"

	"github.com/thebluefowl/tunnely/tunnel"
)

type InMemoryDomainStore map[string]*tunnel.Tunnel

var s *InMemoryDomainStore

func GetInMemoryDomainStore() DomainStore {
	if s != nil {
		return s
	}
	return &InMemoryDomainStore{}
}

func (m *InMemoryDomainStore) GetTunnel(domain string) *tunnel.Tunnel {
	return (*m)[domain]
}

func (m *InMemoryDomainStore) AddTunnel(t *tunnel.Tunnel) {
	fmt.Printf("tunnel created for: %s", t.Subdomain)
	(*m)[t.Subdomain] = t
}
