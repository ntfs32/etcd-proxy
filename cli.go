package main

import (
	"os"
	"github.com/urfave/cli"
	"github.com/ntfs32/etcd-proxy/config"
)

var Config config.ConfigMap

func init()  {
	Config = config.LoadConfig()
}

func main() {
	app := cli.NewApp()
	app.Name = "etcd-proxy"
	app.Usage = "etcd manager cli monitor"
	app.Version = Config.Version
	app.Commands = []cli.Command{
		{
			Name:    "monit",
			Aliases: []string{"m"},
			Usage:   "start a monitor for etcd",
			Action: func(c *cli.Context) error {
				StartMonit()
				return nil
			},
		},
	}
	app.Action=  func(c *cli.Context) error {
		return nil
	}
	app.Run(os.Args)
}


