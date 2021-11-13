package main

import (
	"encoding/json"
	"fmt"
	"gek_net"
	"net"
	"net/http"
)

type Client struct {
	RemoteAddress string `json:"Remote-Address"`
	UserAgent     string `json:"User-Agent"`
}

func httpReturnClientNetworkInfo(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		var client Client
		resp.Header().Set("Content-Type", "application/json")

		// client结构体中填充数据
		client.RemoteAddress = gek_net.GetClientIP(req)
		client.UserAgent = req.Header.Values("User-Agent")[0]

		// 转换为json []byte
		respBody, err := json.Marshal(client)
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

type LocalNetworkInfo struct {
	IPV4  []net.IP
	IPV4P []net.IP
	IPV6  []net.IP
	IPV6P []net.IP
}

func GetLocalNetworkInfo() error {
	var info LocalNetworkInfo
	netInterfaces, err := gek_net.NetInterfaces()
	if err != nil {
		return err
	}
	for _, netInterface := range netInterfaces {
		if t1 := netInterface.GetPublicIP("ipv4"); t1 != nil {
			info.IPV4 = append(info.IPV4, t1)
		}
		if t2 := netInterface.GetPublicIP("ipv6"); t2 != nil {
			info.IPV6 = append(info.IPV4, t2)
		}
		if t3 := netInterface.GetPrivateIP("ipv4"); t3 != nil {
			info.IPV4P = append(info.IPV4, t3)
		}
		if t4 := netInterface.GetPrivateIP("ipv6"); t4 != nil {
			info.IPV6P = append(info.IPV4, t4)
		}
	}

	// 如果全为空则无网络
	if len(info.IPV4)+len(info.IPV4P)+len(info.IPV6)+len(info.IPV6P) == 0 {
		return fmt.Errorf("no network available")
	}

	// 不为空则依次打印
	if len(info.IPV4) != 0 {
		for index, ipv4 := range info.IPV4 {
			fmt.Printf("Public IPV4 %d: %s\n", index+1, ipv4)
		}
	}
	if len(info.IPV4P) != 0 {
		for index, ipv4p := range info.IPV4P {
			fmt.Printf("Private IPV4 %d: %s\n", index+1, ipv4p)
		}
	}
	if len(info.IPV6) != 0 {
		for index, ipv6 := range info.IPV6 {
			fmt.Printf("Public IPV6 %d: %s\n", index+1, ipv6)
		}
	}
	if len(info.IPV6P) != 0 {
		for index, ipv6p := range info.IPV6P {
			fmt.Printf("Private IPV6 %d: %s\n", index+1, ipv6p)
		}
	}

	return nil
}
