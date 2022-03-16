package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	cliAddr    string
	cliPort    int
	cliServe   bool
	cliHelp    bool
	cliVersion bool
)

func init() {
	flag.StringVar(&cliAddr, "addr", "127.0.0.1", "-addr 127.0.0.1")
	flag.IntVar(&cliPort, "port", 1996, "-port 1996")
	flag.BoolVar(&cliServe, "server", false, "-server")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Version:
Usage:
  netinfo {Command} [Option]

Command:
  -server           : start http server
  -h                : show help
  -v                : show version

Option:
  -addr  <IP>       : set server IP
  -port  <Port>     : set server port

Example:
  1) netinfo
  2) netinfo -server
  3) netinfo -server -addr 127.0.0.1 -port 1996
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
		fmt.Println("v1.03")
		os.Exit(0)
	}

	// 如果无 args 返回本地网络信息
	if len(os.Args) == 1 {
		err := GetLocalNetworkInfo()
		if err != nil {
			log.Panicln(err)
		}
		os.Exit(0)
	}
}

func showChangelog() {
	var versionInfo = `Changelog:
  1.00:
    - First release
  1.01:
    - Change GetLocalNetworkInfo() to match the multi-network environment
  1.02:
    - Update use golang 1.17
  1.03:
    - user agent use string instead of []string`

	fmt.Println(versionInfo)
}

func main() {
	// 启动服务
	if cliServe {
		err := startService(cliAddr, cliPort)
		if err != nil {
			log.Panicln(err)
		}
	}
}
