package store

import (
	"fmt"

	"github.com/thebluefowl/tunnely/tunnel"
)

type InMemTunnelStore struct {
	DomainTunnelMap tunnelMap
	IDTunnelMap     tunnelMap
}

type tunnelMap map[string]*tunnel.Tunnel

var inMemTunnelStore *InMemTunnelStore

func GetTunnelStore() TunnelStore {
	// Return global if initialized.
	if inMemTunnelStore != nil {
		return inMemTunnelStore
	}

	// If singleton is not initialized, initialize and return.
	inMemTunnelStore = &InMemTunnelStore{
		DomainTunnelMap: tunnelMap{},
		IDTunnelMap:     tunnelMap{},
	}
	return inMemTunnelStore
}

func (s *InMemTunnelStore) GetTunnelByDomain(domain string) *tunnel.Tunnel {
	return s.DomainTunnelMap[domain]
}

func (s *InMemTunnelStore) GetTunnelByID(id string) *tunnel.Tunnel {
	fmt.Printf("%+v", s)
	return s.IDTunnelMap[id]
}
func (s *InMemTunnelStore) AddTunnel(t *tunnel.Tunnel) {
	s.DomainTunnelMap[t.Subdomain] = t
	s.IDTunnelMap[t.ID] = t
}

func (s *InMemTunnelStore) GetDomains() []string {
	domains := []string{}
	for k := range s.DomainTunnelMap {
		domains = append(domains, k)
	}
	return domains
}

func (s *InMemTunnelStore) GetIDs() []string {
	ids := []string{}
	for k := range s.DomainTunnelMap {
		ids = append(ids, k)
	}
	return ids
}
