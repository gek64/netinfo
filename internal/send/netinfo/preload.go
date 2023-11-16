package netinfo

import (
	"netinfo/internal/receive/controllers/recordService"
	"netinfo/internal/startup"
)

func NewPreload(id string) (preload recordService.RecordBody, err error) {
	netInterfaces, err := startup.GetNetInterfaces()
	if err != nil {
		return recordService.RecordBody{}, err
	}

	preload.ID = id
	preload.NetInterfaces = netInterfaces

	return preload, err
}
