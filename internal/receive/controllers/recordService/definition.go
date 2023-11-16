package recordService

import (
	"netinfo/internal/startup"
)

type RecordBody struct {
	ID            string                 `json:"id" xml:"id" form:"id" binding:"required"`
	NetInterfaces []startup.NetInterface `json:"netInterfaces" xml:"netInterfaces" form:"netInterfaces" binding:"required"`
}

type RecordQuery struct {
	ID string `json:"id" xml:"id" form:"id" binding:"required"`
}
