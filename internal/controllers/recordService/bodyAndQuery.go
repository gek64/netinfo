package recordService

import "netinfo/ent/schema"

type RecordBody struct {
	ID            string                `json:"id" form:"id" binding:"required"`
	Description   string                `json:"description" form:"description"`
	RequestIP     string                `json:"requestIP" form:"requestIP"`
	NetInterfaces []schema.NetInterface `json:"netInterfaces" form:"netInterfaces" binding:"required"`
}

type RecordQuery struct {
	ID string `json:"id" form:"id" binding:"required"`
}
