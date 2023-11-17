package file

import (
	"github.com/gek64/gek/gJson"
	"log"
	"netinfo/internal/send"
	"time"
)

func SendRequest(file string, id string) (err error) {
	// 组装负载
	preload, err := send.NewPreload(id)
	if err != nil {
		return err
	}

	jsonOperator, err := gJson.NewJsonOperator(&preload)
	if err != nil {
		return err
	}

	return jsonOperator.WriteToFile(file)
}

func SendRequestLoop(file string, id string, interval time.Duration) {
	for {
		err := SendRequest(file, id)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("save file to %s\n", file)
		}
		time.Sleep(interval)
	}
}
