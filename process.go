package main

import (
	"fmt"
	"gek_net"
	"net"
	"net/http"
	"strconv"
)

func startService(addr string, port int) (err error) {
	// 检查IP
	if ip := net.ParseIP(addr); ip == nil {
		return fmt.Errorf("%s is not a valid host", addr)
	}
	// 检查端口
	if port < 0 || port > 65535 {
		return fmt.Errorf("%d is not a valid port", port)
	}

	fmt.Println("visit:", addr+":"+strconv.Itoa(port))

	http.HandleFunc("/", httpReturnClientNetworkInfo)
	return http.ListenAndServe(addr+":"+strconv.Itoa(port), nil)
}

func getActiveNetworkInterface() error {
	netInterfaces, err := gek_net.GetNetInterfaces()
	if err != nil {
		return err
	}

	for _, netInterface := range netInterfaces {
		if !netInterface.Flag.Up {
			continue
		}
		fmt.Printf("Active network interface: %s\n", netInterface.Name)
		for _, ipv4 := range netInterface.Ipv4 {
			fmt.Printf("IPv4: %s\n", ipv4)
		}
		for _, ipv6 := range netInterface.Ipv6 {
			fmt.Printf("IPv6: %s\n", ipv6)
		}
	}

	return err
}
