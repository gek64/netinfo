package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/url"
	"netinfo/internal/database"
	"netinfo/internal/ipClient"
	"netinfo/internal/middleware"
	"netinfo/internal/routers"
	"os"
	"time"
)

var (
	cliShow   bool
	cliServer string
	cliClient string

	cliInterval    time.Duration
	cliDescription string

	cliHelp    bool
	cliVersion bool
)

func init() {
	flag.BoolVar(&cliShow, "show", false, "-show")
	flag.StringVar(&cliServer, "server", "", "-server localhost:1996")
	flag.StringVar(&cliClient, "client", "", "-client http://localhost:1996/record")

	flag.DurationVar(&cliInterval, "interval", time.Hour, "-interval 1h")
	flag.StringVar(&cliDescription, "description", "", "-description home_pc")

	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Usage:
	 netinfo {Command} [Option]
	
	Command:
	 -client           : start client
	 -server           : start server
	 -h                : show help
	 -v                : show version
	
	Option:
	 -interval       <IP>     : set client interval
	 -description    <Port>   : set client description
	
	Example:
	 1) netinfo -show
	 2) netinfo -server localhost:1996
	 3) netinfo -client http://localhost:1996/record -interval 1h -description home_opnsense
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
		fmt.Println("v2.00")
		os.Exit(0)
	}

	// 如果无 args 返回本地网络信息
	if len(os.Args) == 1 || cliShow {
		err := ipClient.PrintNetInterfaces()
		if err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}

	if cliServer != "" && cliClient != "" {
		log.Println("Only one of server mode and client mode can be selected")
		os.Exit(0)
	}
}

func sendToServer(url string, interval time.Duration, description string) {
	for {
		_, err := ipClient.Create(url, description)
		if err != nil {
			_, err := ipClient.Update(url, description)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("update completed")
			}
		} else {
			log.Println("created completed")
		}

		time.Sleep(interval)
	}
}

func main() {
	if cliClient != "" {
		targetURL, err := url.Parse(cliClient)
		if err != nil {
			log.Fatalln("invalid client url")
		}
		sendToServer(targetURL.String(), cliInterval, cliDescription)
	}

	if cliServer != "" {
		serverURL, err := url.Parse(cliServer)
		if err != nil {
			log.Fatalln("invalid server url")
		}

		// 创建默认路由引擎,上下文
		engine := gin.Default()
		ctx := context.Background()

		// 初始化数据库
		client, err := database.NewSqliteClient(":memory:")
		if err != nil {
			return
		}
		defer client.Close()

		// 数据库表同步
		err = client.Schema.Create(ctx)
		if err != nil {
			log.Panicf("failed creating schema resources: %v\n", err)
		}

		// 中间件传递参数
		engine.Use(middleware.ParameterPasser(client, ctx))
		// 加载路由
		routers.LoadRouters(engine)
		// 启动
		err = engine.Run(serverURL.String())
		if err != nil {
			log.Println(err)
		}
	}
}
