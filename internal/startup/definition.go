package startup

import "net/netip"

type NetInterface struct {
	Name string       `json:"name"`
	IPs  []netip.Addr `json:"ips"`
	Mac  string       `json:"mac,omitempty"`
}

var Version = "v3.00"

var HelpInfomation = `1234`
