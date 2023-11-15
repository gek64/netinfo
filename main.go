package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"netinfo/internal/receive/routers"
	"netinfo/internal/startup"
	"os"
	"time"
)

var (
	cliServer string
	cliClient string

	cliInterval              time.Duration
	cliDescription           string
	cliUsername              string
	cliPassword              string
	cliSkipCertificateVerify bool

	cliHelp    bool
	cliVersion bool
)

func init() {
	flag.StringVar(&cliServer, "server", "", "-server localhost:1996")
	flag.StringVar(&cliClient, "client", "", "-client http://localhost:1996/record")

	flag.DurationVar(&cliInterval, "interval", 0, "-interval 1h")
	flag.StringVar(&cliDescription, "description", "", "-description home_pc")
	flag.StringVar(&cliUsername, "username", "", "-username bob")
	flag.StringVar(&cliPassword, "password", "", "-password 123456")
	flag.BoolVar(&cliSkipCertificateVerify, "skip-certificate-verify", false, "-skip-certificate-verify")

	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		fmt.Println(startup.HelpInfomation)
	}

	// 打印帮助信息
	if cliHelp {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		fmt.Println(startup.Version)
		os.Exit(0)
	}

	// 如果无 args 返回本地网络信息
	if len(os.Args) == 1 {
		err := startup.PrintNetInterfaces()
		if err != nil {
			os.Exit(1)
		}

		os.Exit(0)
	}

	if cliServer != "" && cliClient != "" {
		log.Println("only one of server mode and client mode can be selected")
		os.Exit(0)
	}
}

func main() {
	//if cliClient != "" {
	//	targetURL, err := url.Parse(cliClient)
	//	if err != nil {
	//		log.Fatalln("invalid client url")
	//	}
	//
	//	if cliInterval != 0 {
	//		ipClient.SendRequestLoop(targetURL.String(), cliInterval, cliDescription, cliUsername, cliPassword, cliSkipCertificateVerify)
	//	} else {
	//		_, err := ipClient.SendRequest(targetURL.String(), cliDescription, cliUsername, cliPassword, cliSkipCertificateVerify)
	//		if err != nil {
	//			log.Println(err)
	//		} else {
	//			log.Println("update completed")
	//		}
	//	}
	//
	//}
	//

	if cliServer != "" {
		// 创建默认路由引擎,上下文
		engine := gin.Default()

		// 加载路由
		routers.LoadRecordRouters(engine)
		routers.LoadDebugRouters(engine)
		// 启动
		err := engine.Run(cliServer)
		if err != nil {
			log.Println(err)
		}
	}
}
