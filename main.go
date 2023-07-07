package main

import (
	"flag"
	"fmt"
	"log"
	"netinfo/internal/ipClient"
	"os"
)

var (
	cliAddr     string
	cliPort     int
	cliServe    bool
	cliUseNetDB bool
	cliHelp     bool
	cliVersion  bool
)

func init() {
	flag.StringVar(&cliAddr, "address", "127.0.0.1", "-address 127.0.0.1")
	flag.IntVar(&cliPort, "port", 1996, "-port 1996")
	flag.BoolVar(&cliServe, "server", false, "-server")
	flag.BoolVar(&cliUseNetDB, "netdb", false, "-netdb")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Usage:
  netinfo {Command} [Option]

Command:
  -server           : start http server
  -h                : show help
  -v                : show version

Option:
  -address <IP>     : set server IP
  -port    <Port>   : set server port
  -netdb            : use net ip database to get ip info

Example:
  1) netinfo
  2) netinfo -server
  3) netinfo -server -address 127.0.0.1 -port 1996 -netdb
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
		fmt.Println("v1.10")
		os.Exit(0)
	}

	// 如果无 args 返回本地网络信息
	//if len(os.Args) == 1 {
	//	os.Exit(0)
	//}
}

func showChangelog() {
	var versionInfo = `Changelog:
  1.00:
    - First release
  1.01:
    - Change getActiveNetworkInterface() to match the multi-network environment
  1.02:
    - Update use golang 1.17
  1.03:
    - user agent use string instead of []string
  1.10:
    - support use ipinfo.io to get ip info`

	fmt.Println(versionInfo)
}

func main() {
	record, err := ipClient.Update("http://127.0.0.1:9999/record", 1, "测试")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(record)

	//// 创建默认路由引擎,上下文
	//engine := gin.Default()
	//ctx := context.Background()
	//
	//// 初始化数据库
	//client, err := database.NewSqliteClient(":memory:")
	//if err != nil {
	//	return
	//}
	//defer client.Close()
	//
	//// 数据库表同步
	//err = client.Schema.Create(ctx)
	//if err != nil {
	//	log.Panicf("failed creating schema resources: %v\n", err)
	//}
	//
	//// 中间件传递参数
	//engine.Use(middleware.ParameterPasser(client, ctx))
	//// 加载路由
	//routers.LoadRouters(engine)
	//// 启动
	//err = engine.Run("127.0.0.1" + ":" + "9999")
	//if err != nil {
	//	log.Panicln(err)
	//}
}
