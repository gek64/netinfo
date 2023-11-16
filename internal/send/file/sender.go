package file

import (
	"github.com/gek64/gek/gJson"
	"log"
	"netinfo/internal/send/netinfo"
	"time"
)

func SendRequest(file string, id string) (err error) {
	// 组装负载
	preload, err := netinfo.NewPreload(id)
	if err != nil {
		return err
	}

	jsonOperator, err := gJson.NewJsonOperator(&preload)
	if err != nil {
		return err
	}

	return jsonOperator.WriteToFile(file)
}

func SendRequestLoop(file string, interval time.Duration, id string) {
	for {
		err := SendRequest(file, id)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("sent successfully using mode file")
		}

		time.Sleep(interval)
	}
}
