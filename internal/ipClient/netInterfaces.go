package ipClient

import (
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"github.com/gek64/gek/gNet"
	"net/netip"
	"netinfo/ent/schema"
	"netinfo/internal/controllers/recordService"
)

func GetNetInterfaces() (netInterfaces []schema.NetInterface, err error) {
	nis, err := gNet.GetNetInterfaces()
	if err != nil {
		return nil, err
	}

	for _, ni := range nis {
		// 只取网路接口标记为UP的
		if !ni.Flag.Up {
			continue
		}

		// 拼接转换网络接口中的IPV4 与IPV6地址
		var ips []netip.Addr
		for _, ipString := range append(ni.Ipv4, ni.Ipv6...) {
			// 跳过回环地址
			if isLoopback, _ := gNet.IsLoopback(ipString); isLoopback {
				continue
			}
			// 跳过地址转换出错
			ipAddr, err := netip.ParseAddr(ipString)
			if err != nil {
				continue
			}

			ips = append(ips, ipAddr)
		}

		// 跳过回环网络接口
		if len(ips) > 0 {
			netInterfaces = append(netInterfaces, schema.NetInterface{
				Name: ni.Name,
				IPs:  ips,
				Mac:  ni.Mac,
			})
		}
	}

	return netInterfaces, nil
}

func GetCreateRecordBody(description string) (createRecordBody recordService.CreateRecordBody, err error) {
	netInterfaces, err := GetNetInterfaces()
	if err != nil {
		return recordService.CreateRecordBody{}, err
	}

	deviceID, err := machineid.ID()
	if err != nil {
		return recordService.CreateRecordBody{}, err
	}

	createRecordBody.ID = deviceID
	createRecordBody.Description = description
	createRecordBody.NetInterfaces = netInterfaces

	return createRecordBody, err
}

func GetUpdateRecordBody(description string) (updateRecordBody recordService.UpdateRecordBody, err error) {
	netInterfaces, err := GetNetInterfaces()
	if err != nil {
		return recordService.UpdateRecordBody{}, err
	}

	deviceID, err := machineid.ID()
	if err != nil {
		return recordService.UpdateRecordBody{}, err
	}

	updateRecordBody.ID = deviceID
	updateRecordBody.Description = description
	updateRecordBody.NetInterfaces = netInterfaces

	return updateRecordBody, err
}

func PrintNetInterfaces() (err error) {
	netInterfaces, err := GetNetInterfaces()
	if err != nil {
		return err
	}
	for i, netInterface := range netInterfaces {
		fmt.Printf("interface: %d\nname: %s\nmac: %s\nips: %v\n", i, netInterface.Name, netInterface.Mac, netInterface.IPs)
	}
	return nil
}
