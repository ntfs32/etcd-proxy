package config

import (
	"os"
	"path/filepath"
	"log"

	"github.com/jinzhu/configor"
)

type ConfigMap struct {
	Backend string `default:"etcd"`
	Version string `default:"1.0.0"`
	Etcd *Etcd
	Etcdv3 *Etcdv3
}

type Etcd struct {
	Endpoints []string `default:"[http://127.0.0.1:2379]"`
	BasePath  string   `default:"/"`
	CaFile    string   `default:"ca_file"`
	CertFile  string   `yaml:"cert_file"`
	KeyFile   string   `yaml:"key_file"`
}

type Etcdv3 struct {
	HostPath []string `default:"[http://127.0.0.1:2379]"`
	Prefix   string `default:"/"`
}

var (
	defaultConfig1 string = "/etc/bdhcfg/server.toml"
	defaultConfig2 string = "./conf/server.toml"

)

func LoadConfig()( config ConfigMap) {
	var (
		filename string
		err error
	)
	if _, err = os.Open(defaultConfig1); err == nil {
		filename = defaultConfig1
	} else {
		//log.Print("no " + defaultConfig1 + " file, use default " + defaultConfig2 + " config ")

		filename, _ = filepath.Abs(defaultConfig2)

		if _, err = os.Open(filename); err != nil {

			log.Print("Error: not any config file,use system default config ", err)
		}
	}
	configor.Load(&config, filename)
	return
}
