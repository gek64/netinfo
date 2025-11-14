package file

import (
	"log"
	"netinfo/internal/send/preload"
	"os"
	"time"
)

func SendRequest(file string, encryptionKey []byte) (err error) {
	// 获取负载
	p, err := preload.GetPreload(encryptionKey)
	if err != nil {
		return err
	}
	return os.WriteFile(file, p, 0644)
}

func SendRequestLoop(file string, encryptionKey []byte, interval time.Duration) {
	for {
		err := SendRequest(file, encryptionKey)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("save file to %s\n", file)
		}
		time.Sleep(interval)
	}
}
