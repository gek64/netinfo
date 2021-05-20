package main

import (
	"encoding/json"
	"localhost/vivycore"
	"net/http"
)

type Client struct {
	RemoteAddress string   `json:"Remote-Address"`
	UserAgent     []string `json:"User-Agent"`
}

func getClientInfo(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		var client Client
		resp.Header().Set("Content-Type", "application/json")

		// client结构体中填充数据
		client.RemoteAddress = vivycore.GetClientIP(req)
		client.UserAgent = req.Header.Values("User-Agent")

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
