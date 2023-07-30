package main

import (
    "context"
    "flag"
    "fmt"
    "github.com/denisbrodbeck/machineid"
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
    cliShowId bool
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
    flag.BoolVar(&cliShowId, "showid", false, "-showid")
    flag.StringVar(&cliServer, "server", "", "-server localhost:1996")
    flag.StringVar(&cliClient, "client", "", "-client http://localhost:1996/record")

    flag.DurationVar(&cliInterval, "interval", time.Hour, "-interval 1h")
    flag.StringVar(&cliDescription, "description", "", "-description home_pc")
    flag.StringVar(&cliUsername, "username", "", "-username bob")
    flag.StringVar(&cliPassword, "password", "", "-password 123456")
    flag.BoolVar(&cliSkipCertificateVerify, "skip-certificate-verify", false, "-skip-certificate-verify")

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
  -showid           : show local machine id
  -h                : show help
  -v                : show version
	
Option:
  -interval      <IP>          : set client interval
  -description   <Port>        : set client description
  -username      <Username>    : set client basic auth username
  -password      <Password>    : set client password
  -skip-certificate-verify     : skip tls certificate verification for http requests
	
Example:
  1) netinfo
  2) netinfo -showid
  3) netinfo -server localhost:1996
  4) netinfo -client http://localhost:1996/record -interval 1h -description main -username bob -password 123456 -skip-certificate-verify
  5) netinfo -h
  6) netinfo -v`

        fmt.Println(helpInfo)
    }

    // 打印帮助信息
    if cliHelp {
        flag.Usage()
        os.Exit(0)
    }

    // 打印版本信息
    if cliVersion {
        fmt.Println("v2.03")
        os.Exit(0)
    }

    // 如果无 args 返回本地网络信息
    if len(os.Args) == 1 {
        err := ipClient.PrintNetInterfaces()
        if err != nil {
            os.Exit(1)
        }
        os.Exit(0)
    }

    if cliShowId {
        id, err := machineid.ID()
        if err != nil {
            log.Panicln(err)
        }
        fmt.Println(id)
        os.Exit(0)
    }

    if cliServer != "" && cliClient != "" {
        log.Println("only one of server mode and client mode can be selected")
        os.Exit(0)
    }
}

func main() {
    if cliClient != "" {
        targetURL, err := url.Parse(cliClient)
        if err != nil {
            log.Fatalln("invalid client url")
        }
        ipClient.SendRequestLoop(targetURL.String(), cliInterval, cliDescription, cliUsername, cliPassword, cliSkipCertificateVerify)
    }

    if cliServer != "" {
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
        err = engine.Run(cliServer)
        if err != nil {
            log.Println(err)
        }
    }
}
