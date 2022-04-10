package main

import (
	"encoding/json"
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

func httpReturnClientNetworkInfo(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		resp.Header().Set("Content-Type", "application/json")

		var response IPInfoMax

		// 转换为json []byte
		respBody, err := json.Marshal(response)
		if err != nil {
			http.Error(resp, "Internal Server Error", http.StatusInternalServerError)
		}

		// 写入resp中
		_, err = resp.Write(respBody)
		if err != nil {
			http.Error(resp, "Internal Server Error", http.StatusInternalServerError)
		}
	default:
		http.Error(resp, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
