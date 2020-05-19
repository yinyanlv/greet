package utils

import (
	"flag"
)

var Root string
var configDir string

func flags() {

	flag.StringVar(&Root, "root", "", "App root directory")
	flag.StringVar(&configDir, "config", "", "App config files directory")
	flag.Parse()
}