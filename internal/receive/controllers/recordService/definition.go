package recordService

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

type RecordBody struct {
	ID            string                 `json:"id" xml:"id" form:"id" binding:"required"`
	NetInterfaces []startup.NetInterface `json:"netInterfaces" xml:"netInterfaces" form:"netInterfaces" binding:"required"`
}

type RecordQuery struct {
	ID string `json:"id" xml:"id" form:"id" binding:"required"`
}
