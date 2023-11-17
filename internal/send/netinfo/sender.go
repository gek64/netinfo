package netinfo

import (
	"github.com/imroc/req/v3"
	"log"
	"netinfo/internal/send"
	"time"
)

func SendRequest(endpoint string, username string, password string, skipCertVerify bool, id string) (resp *req.Response, err error) {
	client := req.C()

	// 跳过 TLS 证书检测
	if skipCertVerify {
		client.EnableInsecureSkipVerify()
	}

	// 组装负载
	preload, err := send.NewPreload(id)
	if err != nil {
		return nil, err
	}

	// 发送 PUT 请求
	resp, err = client.R().
		SetBody(preload).
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

func SendRequestLoop(endpoint string, username string, password string, skipCertVerify bool, id string, interval time.Duration) {
	for {
		resp, err := SendRequest(endpoint, username, password, skipCertVerify, id)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(resp.Status)
		}
		time.Sleep(interval)
	}
}
