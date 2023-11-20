package netinfo

import (
	"net/netip"
	"time"
)

type Data struct {
	ID            string         `json:"id"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	RequestIP     netip.Addr     `json:"requestIP"`
	NetInterfaces []NetInterface `json:"netInterfaces"`
}

type NetInterface struct {
	Name string       `json:"name"`
	IPs  []netip.Addr `json:"ips"`
	Mac  string       `json:"mac,omitempty"`
}
