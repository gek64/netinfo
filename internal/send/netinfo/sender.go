package netinfo

import (
	"github.com/imroc/req/v3"
	"log"
	"netinfo/internal/send/preload"
	"time"
)

func SendRequest(endpoint string, username string, password string, allowInsecure bool, id string) (resp *req.Response, err error) {
	client := req.C()

	// 跳过 TLS 证书检测
	if allowInsecure {
		client.EnableInsecureSkipVerify()
	}

	// 获取负载
	p, err := preload.GetPreload(id, nil)
	if err != nil {
		return nil, err
	}

	// 发送 PUT 请求
	resp, err = client.R().
		SetBody(p).
		SetRetryCount(3).
		SetRetryBackoffInterval(1*time.Second, 5*time.Second).
		SetBasicAuth(username, password).
		Put(endpoint)
	if err == nil && resp.IsSuccessState() {
		return resp, nil
	} else {
		return resp, err
	}
}

func SendRequestLoop(endpoint string, username string, password string, allowInsecure bool, id string, interval time.Duration) {
	for {
		resp, err := SendRequest(endpoint, username, password, allowInsecure, id)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(resp.Status)
		}
		time.Sleep(interval)
	}
}
