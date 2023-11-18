package webdav

import (
	"github.com/gek64/gek/gWebDAV"
	"log"
	"net/http"
	"netinfo/internal/send/preload"
	"time"
)

func SendRequest(endpoint string, username string, password string, allowInsecure bool, filepath string, id string, encryptionKey []byte) (resp *http.Response, err error) {
	// 获取负载
	p, err := preload.GetPreload(id, encryptionKey)
	if err != nil {
		return nil, err
	}

	client, err := gWebDAV.NewClient(endpoint, username, password, allowInsecure)
	if err != nil {
		return nil, err
	}

	return client.Upload(filepath, p)
}

func SendRequestLoop(endpoint string, username string, password string, allowInsecure bool, filepath string, id string, encryptionKey []byte, interval time.Duration) {
	for {
		resp, err := SendRequest(endpoint, username, password, allowInsecure, filepath, id, encryptionKey)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(resp.Status)
		}
		time.Sleep(interval)
	}
}
