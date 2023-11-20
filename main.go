package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"log"
	"netinfo/internal/netinfo"
	"netinfo/internal/receive/routers"
	"netinfo/internal/send/file"
	"netinfo/internal/send/nconnect"
	"netinfo/internal/send/s3"
	"netinfo/internal/send/webdav"
	"os"
	"time"
)

func main() {
	// send mode
	var id string
	var allow_insecure bool
	var encryption_key string
	var interval time.Duration
	var endpoint string
	var username string
	var password string

	// send mode file
	var filepath string

	// send mode s3
	var regin string
	var sts_token string
	var path_style bool
	var bucket string
	var object_path string

	// receive mode
	var listen_address string

	cmds := []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list network information",
			Action: func(ctx *cli.Context) error {
				return netinfo.PrintNetInterfaces()
			},
		},
		{
			Name:    "send",
			Aliases: []string{"s"},
			Usage:   "send network information",

			Subcommands: []*cli.Command{
				{
					Name:  "file",
					Usage: "send to filesystem",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:        "id",
							Usage:       "set id",
							Required:    true,
							Destination: &id,
						},
						&cli.StringFlag{
							Name:        "filepath",
							Usage:       "set file path",
							Required:    true,
							Destination: &filepath,
						},
						&cli.StringFlag{
							Name:        "encryption_key",
							Usage:       "set file encryption key",
							Destination: &encryption_key,
						},
						&cli.DurationFlag{
							Name:        "interval",
							Usage:       "set send interval",
							Destination: &interval,
						},
					},
					Action: func(ctx *cli.Context) error {
						if interval != 0 {
							file.SendRequestLoop(filepath, id, []byte(encryption_key), interval)
						} else {
							return file.SendRequest(filepath, id, []byte(encryption_key))
						}
						return nil
					},
				},
				{
					Name:  "s3",
					Usage: "send to s3 server",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:        "id",
							Usage:       "set id",
							Required:    true,
							Destination: &id,
						},
						&cli.BoolFlag{
							Name:        "allow_insecure",
							Usage:       "set allow insecure connect",
							Value:       false,
							Destination: &allow_insecure,
						},
						&cli.StringFlag{
							Name:        "encryption_key",
							Usage:       "set file encryption key",
							Destination: &encryption_key,
						},
						&cli.DurationFlag{
							Name:        "interval",
							Usage:       "set send interval",
							Destination: &interval,
						},
						&cli.StringFlag{
							Name:        "endpoint",
							Usage:       "set s3 server endpoint",
							Required:    true,
							Destination: &endpoint,
						},
						&cli.StringFlag{
							Name:        "regin",
							Usage:       "set s3 server regin",
							Value:       "us-east-1",
							Destination: &regin,
						},
						&cli.StringFlag{
							Name:        "access_key_id",
							Usage:       "set s3 server access key id",
							Required:    true,
							Destination: &username,
						},
						&cli.StringFlag{
							Name:        "secret_access_key",
							Usage:       "set s3 server secret access key",
							Required:    true,
							Destination: &password,
						},
						&cli.StringFlag{
							Name:        "sts_token",
							Usage:       "set s3 server sts token",
							Destination: &sts_token,
						},
						&cli.BoolFlag{
							Name:        "path_style",
							Usage:       "set s3 server path style, false: virtual host, true: path",
							Value:       false,
							Destination: &path_style,
						},
						&cli.StringFlag{
							Name:        "bucket",
							Usage:       "set s3 server bucket",
							Required:    true,
							Destination: &bucket,
						},
						&cli.StringFlag{
							Name:        "object_path",
							Usage:       "set s3 server object path",
							Required:    true,
							Destination: &object_path,
						},
					},
					Action: func(ctx *cli.Context) error {
						if interval != 0 {
							s3.SendRequestLoop(endpoint, regin, username, password, sts_token, path_style, allow_insecure, bucket, object_path, id, []byte(encryption_key), interval)
						} else {
							_, err := s3.SendRequest(endpoint, regin, username, password, sts_token, path_style, allow_insecure, bucket, object_path, id, []byte(encryption_key))
							if err != nil {
								return err
							}
						}
						return nil
					},
				},
				{
					Name:  "webdav",
					Usage: "send to webdav server",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:        "id",
							Usage:       "set id",
							Required:    true,
							Destination: &id,
						},
						&cli.BoolFlag{
							Name:        "allow_insecure",
							Usage:       "set allow insecure connect",
							Value:       false,
							Destination: &allow_insecure,
						},
						&cli.StringFlag{
							Name:        "encryption_key",
							Usage:       "set file encryption key",
							Destination: &encryption_key,
						},
						&cli.DurationFlag{
							Name:        "interval",
							Usage:       "set send interval",
							Destination: &interval,
						},
						&cli.StringFlag{
							Name:        "endpoint",
							Usage:       "set webdav server endpoint",
							Required:    true,
							Destination: &endpoint,
						},
						&cli.StringFlag{
							Name:        "username",
							Usage:       "set webdav server username",
							Destination: &username,
						},
						&cli.StringFlag{
							Name:        "password",
							Usage:       "set webdav server password",
							Destination: &password,
						},
						&cli.StringFlag{
							Name:        "filepath",
							Usage:       "set webdav server filepath",
							Required:    true,
							Destination: &filepath,
						},
					},
					Action: func(ctx *cli.Context) error {
						if interval != 0 {
							webdav.SendRequestLoop(endpoint, username, password, allow_insecure, filepath, id, []byte(encryption_key), interval)
						} else {
							_, err := webdav.SendRequest(endpoint, username, password, allow_insecure, filepath, id, []byte(encryption_key))
							if err != nil {
								return err
							}
						}
						return nil
					},
				},
				{
					Name:  "nconnect",
					Usage: "send to nconnect server",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:        "id",
							Usage:       "set id",
							Required:    true,
							Destination: &id,
						},
						&cli.BoolFlag{
							Name:        "allow_insecure",
							Usage:       "set allow insecure connect",
							Value:       false,
							Destination: &allow_insecure,
						},
						&cli.DurationFlag{
							Name:        "interval",
							Usage:       "set send interval",
							Destination: &interval,
						},
						&cli.StringFlag{
							Name:        "endpoint",
							Usage:       "set nconnect server endpoint",
							Required:    true,
							Destination: &endpoint,
						},
					},
					Action: func(ctx *cli.Context) error {
						if interval != 0 {
							nconnect.SendRequestLoop(endpoint, username, password, allow_insecure, id, interval)
						} else {
							_, err := nconnect.SendRequest(endpoint, username, password, allow_insecure, id)
							if err != nil {
								return err
							}
						}
						return nil
					},
				},
			},
		},
		{
			Name:    "receive",
			Aliases: []string{"r"},
			Usage:   "receive network information",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "listen",
					Aliases:     []string{"l"},
					Usage:       "set nconnect server listen address",
					Value:       "127.0.0.1:1996",
					Destination: &listen_address,
				},
			},
			Action: func(context *cli.Context) error {
				// 创建默认路由引擎
				engine := gin.Default()
				routers.LoadRecordRouters(engine)
				routers.LoadDebugRouters(engine)

				// 启动
				return engine.Run(listen_address)
			},
		},
	}

	// 打印版本函数
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("%s", cCtx.App.Version)
	}

	app := &cli.App{
		Usage:    "Network information manager",
		Version:  "v3.00",
		Commands: cmds,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
