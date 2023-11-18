package file

import (
	"log"
	"netinfo/internal/send/preload"
	"os"
	"time"
)

func SendRequest(file string, id string, encryptedKey []byte) (err error) {
	// 获取负载
	p, err := preload.GetPreload(id, encryptedKey)
	if err != nil {
		return err
	}
	return os.WriteFile(file, p, 0644)
}

func SendRequestLoop(file string, id string, encryptedKey []byte, interval time.Duration) {
	for {
		err := SendRequest(file, id, encryptedKey)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("save file to %s\n", file)
		}
		time.Sleep(interval)
	}
}
