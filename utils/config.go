package utils

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"path"
	"sync"
)

var (
	once sync.Once
	Cfg  Config
)

type Config struct {
	Name string `toml:"name"`
}

func init() {
	flags()
 	initCfg()
}

func initCfg() {

	once.Do(func() {
		p := configPath()

		if _, err := toml.DecodeFile(p, &Cfg); err != nil {
			fmt.Println(err)
		}
	})
}

func configPath() string {
	return path.Join("./", "/config/config.toml")
}
