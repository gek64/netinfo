package netinfo

import (
	"github.com/imroc/req/v3"
	"log"
	"time"
)

func SendRequest(url string, id string, username string, password string, skipCertVerify bool) (resp *req.Response, err error) {
	client := req.C()

	// 跳过 TLS 证书检测
	if skipCertVerify {
		client.EnableInsecureSkipVerify()
	}

	// 组装负载
	preload, err := NewPreload(id)
	if err != nil {
		return nil, err
	}

	// 发送 PUT 请求
	resp, err = client.R().
		SetBody(preload).
		SetRetryCount(3).
		SetRetryBackoffInterval(1*time.Second, 5*time.Second).
		SetBasicAuth(username, password).
		Put(url)
	if err == nil && resp.IsSuccessState() {
		return resp, nil
	} else {
		return resp, err
	}
}

func SendRequestLoop(url string, interval time.Duration, id string, username string, password string, skipCertVerify bool) {
	for {
		_, err := SendRequest(url, id, username, password, skipCertVerify)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("sent successfully using mode netinfo")
		}

		time.Sleep(interval)
	}
}
