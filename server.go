package main

import (
	"fmt"
	"gek_net"
	"log"
	"net"
)

func getServerInfo() {
	var IPv4, IPv6, IPv4p, IPv6p net.IP
	netInterfaces, err := gek_net.NetInterfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, netInterface := range netInterfaces {
		if t1 := netInterface.GetPublicIP("ipv4"); t1 != nil {
			IPv4 = t1
		}
		if t2 := netInterface.GetPublicIP("ipv6"); t2 != nil {
			IPv6 = t2
		}
		if t3 := netInterface.GetPrivateIP("ipv4"); t3 != nil {
			IPv4p = t3
		}
		if t4 := netInterface.GetPrivateIP("ipv6"); t4 != nil {
			IPv6p = t4
		}
	}

	if IPv4 == nil && IPv6 == nil && IPv4p == nil && IPv6p == nil {
		fmt.Println("No Network Available")
	} else {
		if IPv4 != nil {
			fmt.Println("Server Public IPV4:", IPv4)
		}
		if IPv6 != nil {
			fmt.Println("Server Public IPV6:", IPv6)
		}
		if IPv4p != nil {
			fmt.Println("Server Private IPV4:", IPv4p)
		}
		if IPv6p != nil {
			fmt.Println("Server Private IPV6:", IPv6p)
		}
	}
}
