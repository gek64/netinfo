package netinfo

import (
	"fmt"
	"github.com/gek64/gek/gNet"
	"net/netip"
)

func GetNetInterfaces() (netInterfaces []NetInterface, err error) {
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
			// 回环地址
			isLoopback, _ := gNet.IsLoopback(ipString)
			if isLoopback {
				continue
			}
			// 链路本地地址
			isLinkLocal, _ := gNet.IsLinkLocal(ipString)
			if isLinkLocal {
				continue
			}
			// 专用网络地址
			isPrivate, _ := gNet.IsPrivate(ipString)
			if isPrivate {
				//continue
			}
			// 地址转换
			ipAddr, err := netip.ParseAddr(ipString)
			if err != nil {
				continue
			}
			ips = append(ips, ipAddr)
		}

		// 跳过回环网络接口
		if len(ips) > 0 {
			netInterfaces = append(netInterfaces, NetInterface{
				Name: ni.Name,
				IPs:  ips,
				Mac:  ni.Mac,
			})
		}
	}

	return netInterfaces, nil
}

func PrintNetInterfaces() (err error) {
	netInterfaces, err := GetNetInterfaces()
	if err != nil {
		return err
	}
	for _, netInterface := range netInterfaces {
		fmt.Printf("name: %s\nmac: %s\nips: %v\n", netInterface.Name, netInterface.Mac, netInterface.IPs)
	}
	return nil
}
