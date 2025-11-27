package webdav

import (
	"log"
	"net/http"
	"netinfo/internal/send/preload"
	"time"

	"github.com/unix755/xtools/xWebDAV"
)

func SendRequest(endpoint string, username string, password string, allowInsecure bool, filepath string, encryptionKey []byte) (resp *http.Response, err error) {
	// 获取负载
	p, err := preload.GetPreload(encryptionKey)
	if err != nil {
		return nil, err
	}

	client, err := xWebDAV.NewClient(endpoint, username, password, allowInsecure)
	if err != nil {
		return nil, err
	}

	return client.Upload(filepath, p)
}

func SendRequestLoop(endpoint string, username string, password string, allowInsecure bool, filepath string, encryptionKey []byte, interval time.Duration) {
	for {
		resp, err := SendRequest(endpoint, username, password, allowInsecure, filepath, encryptionKey)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(resp.Status)
		}
		time.Sleep(interval)
	}
}
