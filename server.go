package main

import (
	"encoding/json"
	"github.com/gek64/gek/gNet"
	"net/http"
)

func httpIpInfo(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		resp.Header().Set("Content-Type", "application/json")

		// 获取resp body
		respBody, err := respBodyBuilder(req, req.Header.Values("User-Agent")[0], cliUseNetDB)
		if err != nil {
			http.Error(resp, "Internal Server Error, Fail to build response body", http.StatusInternalServerError)
		}

		// 写入resp
		_, err = resp.Write(respBody)
		if err != nil {
			http.Error(resp, "Internal Server Error, Fail to write data to response", http.StatusInternalServerError)
		}

	default:
		http.Error(resp, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func respBodyBuilder(req *http.Request, useragent string, useNetDB bool) (respBody []byte, err error) {
	// 获取req中真实ip
	ip, err := gNet.GetIPFromRequest(req)
	if err != nil {
		return nil, err
	}

	// 如果要使用网络ip数据库
	if useNetDB {
		ipInfo, err := getIPInfoFromIpInfoIo(ip)
		if err == nil {
			respData := Resp{IPInfo: ipInfo}
			respData.UserAgent = useragent
			// 转换为json []byte
			respBody, err = json.Marshal(respData)
			if err != nil {
				return nil, err
			}
			return respBody, nil
		}
	}

	// 不使用网络ip数据库,或者连接网络ip数据库失败
	respData := RespMin{IP: ip}
	respData.UserAgent = useragent
	// 转换为json []byte
	respBody, err = json.Marshal(respData)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
