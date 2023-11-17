package s3

import (
	"encoding/json"
	"github.com/gek64/gek/gS3"
	"log"
	"netinfo/internal/send"
	"time"
)

func SendRequest(endpoint string, region string, accessKeyId string, secretAccessKey string, stsToken string, pathStyle bool, allowInsecure bool, bucket string, filename string, id string) (location *string, err error) {
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

	// 使用 s3 协议上传负载
	s := gS3.NewS3Session(endpoint, region, accessKeyId, secretAccessKey, stsToken, pathStyle, allowInsecure)
	result, err := s.UploadObject(bucket, filename, preloadBytes)
	if err != nil {
		return &result.Location, err
	}
	return nil, nil
}

func SendRequestLoop(endpoint string, region string, accessKeyId string, secretAccessKey string, stsToken string, pathStyle bool, allowInsecure bool, bucket string, filename string, id string, interval time.Duration) {
	for {
		location, err := SendRequest(endpoint, region, accessKeyId, secretAccessKey, stsToken, pathStyle, allowInsecure, bucket, filename, id)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("upload to %s", *location)
		}
		time.Sleep(interval)
	}
}
