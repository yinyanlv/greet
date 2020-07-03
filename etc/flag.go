package etc

import (
	"flag"
)

var AppMode string

func flags() {
	flag.StringVar(&AppMode, "mode", "dev", "App run mode")
	flag.Parse()
}