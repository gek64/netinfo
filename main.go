package main

import (
	"encoding/json"
	"fmt"
	"log"
	"netinfo/internal/netinfo"
	"netinfo/internal/send/file"
	"netinfo/internal/send/preload"
	"netinfo/internal/send/s3"
	"netinfo/internal/send/webdav"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	// show mode
	var showPreload bool

	// send mode
	var allowInsecure bool
	var encryptionKey string
	var interval time.Duration
	var endpoint string
	var username string
	var password string

	// send mode file
	var filepath string

	// send mode s3
	var regin string
	var stsToken string
	var pathStyle bool
	var bucket string
	var objectPath string

	cmds := []*cli.Command{
		{
			Name:  "show",
			Usage: "show all network information",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "preload",
					Aliases:     []string{"p"},
					Usage:       "show preload information",
					Value:       false,
					Destination: &showPreload,
				},
			},
			Action: func(ctx *cli.Context) error {
				var err error
				var p []byte
				var netInterfaces []netinfo.NetInterface

				if showPreload {
					p, err = preload.GetPreload([]byte(encryptionKey))
				} else {
					netInterfaces, err = netinfo.GetNetInterfaces()
					p, err = json.Marshal(netInterfaces)
				}

				fmt.Println(string(p))
				return err
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
							Name:        "filepath",
							Usage:       "set file path",
							Required:    true,
							Destination: &filepath,
						},
						&cli.StringFlag{
							Name:        "encryption_key",
							Usage:       "set file encryption key",
							Destination: &encryptionKey,
						},
						&cli.DurationFlag{
							Name:        "interval",
							Usage:       "set send interval",
							Destination: &interval,
						},
					},
					Action: func(ctx *cli.Context) error {
						if interval != 0 {
							file.SendRequestLoop(filepath, []byte(encryptionKey), interval)
						} else {
							return file.SendRequest(filepath, []byte(encryptionKey))
						}
						return nil
					},
				},
				{
					Name:  "s3",
					Usage: "send to s3 server",
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:        "allow_insecure",
							Usage:       "set allow insecure connect",
							Value:       false,
							Destination: &allowInsecure,
						},
						&cli.StringFlag{
							Name:        "encryption_key",
							Usage:       "set file encryption key",
							Destination: &encryptionKey,
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
							Destination: &stsToken,
						},
						&cli.BoolFlag{
							Name:        "path_style",
							Usage:       "set s3 server path style, false: virtual host, true: path",
							Value:       false,
							Destination: &pathStyle,
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
							Destination: &objectPath,
						},
					},
					Action: func(ctx *cli.Context) error {
						if interval != 0 {
							s3.SendRequestLoop(endpoint, regin, username, password, stsToken, pathStyle, allowInsecure, bucket, objectPath, []byte(encryptionKey), interval)
						} else {
							_, err := s3.SendRequest(endpoint, regin, username, password, stsToken, pathStyle, allowInsecure, bucket, objectPath, []byte(encryptionKey))
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
						&cli.BoolFlag{
							Name:        "allow_insecure",
							Usage:       "set allow insecure connect",
							Value:       false,
							Destination: &allowInsecure,
						},
						&cli.StringFlag{
							Name:        "encryption_key",
							Usage:       "set file encryption key",
							Destination: &encryptionKey,
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
							webdav.SendRequestLoop(endpoint, username, password, allowInsecure, filepath, []byte(encryptionKey), interval)
						} else {
							_, err := webdav.SendRequest(endpoint, username, password, allowInsecure, filepath, []byte(encryptionKey))
							if err != nil {
								return err
							}
						}
						return nil
					},
				},
			},
		},
	}

	// 打印版本函数
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("%s", cCtx.App.Version)
	}

	app := &cli.App{
		Usage:    "Network information manager",
		Version:  "v3.10",
		Commands: cmds,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
