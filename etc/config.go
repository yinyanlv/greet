package etc

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"path"
	"sync"
)

var (
	once      sync.Once
	AppConfig Config
)

type Config struct {
	Name       string `toml:"name"`
	Version    string `toml:"version"`
	Mode       string `toml:"mode"`
	Protocol   string `toml:"protocol"`
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	PublicPath string `toml:"public_path"`
	MySQL      MySQLConfig
	Log        LogConfig
}

type MySQLConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DB       string `toml:"db"`
}

type LogConfig struct {
	Dir string `toml:"dir"`
}

func init() {
	flags()
	initCfg()
}

func initCfg() {
	once.Do(func() {
		p := configPath()
		if _, err := toml.DecodeFile(p, &AppConfig); err != nil {
			fmt.Println(err)
		}
	})
}

func configPath() string {
	var p string
	if AppMode != "dev" {
		p = "/config/config_" + AppMode + ".toml"
	} else {
		p = "/config/config.toml"
	}
	return path.Join("./", p)
}
