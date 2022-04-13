package main

import (
	"encoding/json"
	"io"
	"net/http"
)

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
