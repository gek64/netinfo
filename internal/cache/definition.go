package cache

import (
	"net/netip"
	"netinfo/internal/startup"
	"time"
)

const (
	Database = "database"
)

type NetInfoInMemoryData struct {
	ID            string                 `json:"id"`
	UpdatedAt     time.Time              `json:"updatedAt"`
	RequestIP     netip.Addr             `json:"requestIP"`
	NetInterfaces []startup.NetInterface `json:"netInterfaces"`
}
