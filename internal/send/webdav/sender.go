package webdav

import (
	"github.com/gek64/gek/gWebDAV"
	"log"
	"net/http"
	"netinfo/internal/send/preload"
	"time"
)

func SendRequest(endpoint string, filepath string, username string, password string, allowInsecure bool, id string, encryptedKey []byte) (resp *http.Response, err error) {
	// 获取负载
	p, err := preload.GetPreload(id, encryptedKey)
	if err != nil {
		return nil, err
	}

	client, err := gWebDAV.NewClient(endpoint, username, password, allowInsecure)
	if err != nil {
		return nil, err
	}

	return client.Upload(filepath, p)
}

func SendRequestLoop(endpoint string, filepath string, username string, password string, allowInsecure bool, id string, encryptedKey []byte, interval time.Duration) {
	for {
		resp, err := SendRequest(endpoint, filepath, username, password, allowInsecure, id, encryptedKey)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(resp.Status)
		}
		time.Sleep(interval)
	}
}
