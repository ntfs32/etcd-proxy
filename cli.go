package main

import (
	"os"
	"gopkg.in/urfave/cli.v2"
	"github.com/ntfs32/etcd-proxy/config"
)

var Config config.ConfigMap

func init()  {
	Config = config.LoadConfig()
}

func main() {
	app := &cli.App{
		Name:    "etcd-proxy",
		Usage:   "Discovery Config Manager",
		Version: Config.Version,
	}
	app.Action=  func(c *cli.Context) error {

		if (c.Args().Get(0) == "monit"){
			StartMonit()
		}
		return nil
	}
	app.Run(os.Args)
}


