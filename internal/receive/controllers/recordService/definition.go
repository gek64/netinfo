package recordService

import (
	"netinfo/internal/netinfo"
)

const (
	Database = "database"
)

type RecordBody struct {
	ID            string                 `json:"id" xml:"id" form:"id" binding:"required"`
	NetInterfaces []netinfo.NetInterface `json:"netInterfaces" xml:"netInterfaces" form:"netInterfaces" binding:"required"`
}

type RecordQuery struct {
	ID string `json:"id" xml:"id" form:"id" binding:"required"`
}
