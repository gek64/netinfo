package main

import (
	"fmt"
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

	http.HandleFunc("/", httpIpInfo)

	return http.ListenAndServe(addr+":"+strconv.Itoa(port), nil)
}
