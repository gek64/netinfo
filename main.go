package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
)

func main() {
	var host string
	var port int
	var serve bool

	// 设置flag
	flag.StringVar(&host, "h", "127.0.0.1", "-h 127.0.0.1")
	if ip := net.ParseIP(host); ip == nil {
		log.Fatalf("%s is not a valid host", host)
	}
	flag.IntVar(&port, "p", 6000, "-p 6000")
	if port < 0 || port > 65535 {
		log.Fatalf("%d is not a valid port", port)
	}
	flag.BoolVar(&serve, "s", false, "-s")
	flag.Parse()

	// 获取服务器信息
	getServerInfo()

	// 启动服务
	if serve {
		http.HandleFunc("/", getClientInfo)
		fmt.Println("Info Program Serve at:", host+":"+strconv.Itoa(port))
		err := http.ListenAndServe(host+":"+strconv.Itoa(port), nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
