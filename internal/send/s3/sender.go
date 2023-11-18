package s3

import (
	"github.com/gek64/gek/gS3"
	"log"
	"netinfo/internal/send/preload"
	"time"
)

func SendRequest(endpoint string, region string, accessKeyId string, secretAccessKey string, stsToken string, pathStyle bool, allowInsecure bool, bucket string, filename string, id string, encryptedKey []byte) (location *string, err error) {
	// 获取负载
	p, err := preload.GetPreload(id, encryptedKey)
	if err != nil {
		return nil, err
	}

	// 使用 s3 协议上传负载
	s := gS3.NewS3Session(endpoint, region, accessKeyId, secretAccessKey, stsToken, pathStyle, allowInsecure)
	result, err := s.UploadObject(bucket, filename, p)
	if err != nil {
		return nil, err
	}
	return &result.Location, nil
}

func SendRequestLoop(endpoint string, region string, accessKeyId string, secretAccessKey string, stsToken string, pathStyle bool, allowInsecure bool, bucket string, filename string, id string, encryptedKey []byte, interval time.Duration) {
	for {
		location, err := SendRequest(endpoint, region, accessKeyId, secretAccessKey, stsToken, pathStyle, allowInsecure, bucket, filename, id, encryptedKey)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("upload to %s", *location)
		}
		time.Sleep(interval)
	}
}
