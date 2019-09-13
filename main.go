package main

import (
	"os"

	"github.com/ehazlett/gatekeeper/version"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = version.Name
	app.Author = "@darknet"
	app.Description = version.Description
	app.Version = version.BuildVersion()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "enable debug logging",
		},
		cli.IntFlag{
			Name:  "port, p",
			Usage: "listen port",
			Value: 2222,
		},
		cli.StringFlag{
			Name:  "key-dir, d",
			Usage: "path to authorized public keys",
			Value: "",
		},
		cli.StringFlag{
			Name:  "host-key, k",
			Usage: "path to host key",
			Value: "/etc/ssh/ssh_host_rsa_key",
		},
	}
	app.Before = func(cx *cli.Context) error {
		if cx.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}

		return nil
	}
	app.Action = func(cx *cli.Context) error {
		cfg := &ServerConfig{
			ListenPort:  cx.Int("port"),
			KeysPath:    cx.String("key-dir"),
			HostKeyPath: cx.String("host-key"),
		}
		srv, err := NewServer(cfg)
		if err != nil {
			return err
		}
		return srv.Run()
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
