package recordService

import (
	"net/netip"
	"time"
)

const (
	Database = "database"
)

type NetInfoInMemoryData struct {
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

type RecordBody struct {
	ID            string         `json:"id" xml:"id" form:"id" binding:"required"`
	NetInterfaces []NetInterface `json:"netInterfaces" xml:"netInterfaces" form:"netInterfaces" binding:"required"`
}

type RecordQuery struct {
	ID string `json:"id" xml:"id" form:"id" binding:"required"`
}
