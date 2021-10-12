package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

var (
	cliAddr    string
	cliPort    int
	cliServe   bool
	cliHelp    bool
	cliVersion bool
)

func init() {
	flag.StringVar(&cliAddr, "a", "127.0.0.1", "-a 127.0.0.1")
	flag.IntVar(&cliPort, "p", 1996, "-p 1996")
	flag.BoolVar(&cliServe, "s", false, "-s")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 检查IP
	if ip := net.ParseIP(cliAddr); ip == nil {
		log.Fatalf("%s is not a valid host", cliAddr)
	}
	// 检查端口
	if cliPort < 0 || cliPort > 65535 {
		log.Fatalf("%d is not a valid port", cliPort)
	}

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Version:
  1.02

Usage:
  netinfo {Command} [Option]

Command:
  -s                : start http server
  -h                : show help
  -v                : show version

Option:
  -a  <IP>          : set server IP
  -p  <Port>        : set server port

Example:
  1) netinfo
  2) netinfo -s
  3) netinfo -s -a 127.0.0.1 -p 1996
  4) netinfo -h
  5) netinfo -v`

		fmt.Println(helpInfo)
	}

	// 打印帮助信息
	if cliHelp {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		showVersion()
		os.Exit(0)
	}

	// 获取服务器信息
	err := GetLocalNetworkInfo()
	if err != nil {
		log.Fatal(err)
	}
}

func showVersion() {
	var versionInfo = `Changelog:
  1.00:
    - First release
  1.01:
    - Change GetLocalNetworkInfo() to match the multi-network environment
  1.02:
    - Update use golang 1.17`

	fmt.Println(versionInfo)
}

func main() {
	// 启动服务
	if cliServe {
		http.HandleFunc("/", httpReturnClientNetworkInfo)
		fmt.Println("Info Program Serve at:", cliAddr+":"+strconv.Itoa(cliPort))
		err := http.ListenAndServe(cliAddr+":"+strconv.Itoa(cliPort), nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
