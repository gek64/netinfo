package models

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
)

type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
}

type Resp struct {
	IP        net.Addr
	IPInfo    IPInfo
	UserAgent string
}

func getIPInfoFromIpInfoIo(ip string) (ipInfo IPInfo, err error) {
	response, err := http.Get("https://ipinfo.io/" + ip + "/json")
	if err != nil {
		return IPInfo{}, err
	}

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return IPInfo{}, err
	}

	err = json.Unmarshal(respBytes, &ipInfo)
	if err != nil {
		return IPInfo{}, err
	}

	return ipInfo, err
}
