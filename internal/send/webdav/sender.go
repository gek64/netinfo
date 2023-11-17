package webdav

import (
	"encoding/json"
	"github.com/gek64/gek/gWebDAV"
	"log"
	"net/http"
	"netinfo/internal/send"
	"time"
)

func SendRequest(endpoint string, file string, username string, password string, skipCertVerify bool, id string) (resp *http.Response, err error) {
	client, err := gWebDAV.NewClient(endpoint, username, password, skipCertVerify)
	if err != nil {
		return nil, err
	}

	// 组装负载
	preload, err := send.NewPreload(id)
	if err != nil {
		return nil, err
	}
	// 负载结构体转换为比特切片
	preloadBytes, err := json.Marshal(preload)
	if err != nil {
		return nil, err
	}

	return client.Upload(file, preloadBytes)
}

func SendRequestLoop(endpoint string, file string, username string, password string, skipCertVerify bool, id string, interval time.Duration) {
	for {
		resp, err := SendRequest(endpoint, file, username, password, skipCertVerify, id)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(resp.Status)
		}
		time.Sleep(interval)
	}
}
