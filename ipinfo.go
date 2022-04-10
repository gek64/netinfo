package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func getIPInfoFromInternet(ip string) (ipInfo IPInfoMax, err error) {
	response, err := http.Get("https://ipinfo.io/" + ip + "/json")
	if err != nil {
		return IPInfoMax{}, err
	}

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return IPInfoMax{}, err
	}

	err = json.Unmarshal(respBytes, &ipInfo)
	if err != nil {
		return IPInfoMax{}, err
	}

	return ipInfo, err
}
