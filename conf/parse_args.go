package conf

import (
	"fmt"
	"os"
)

// go build -v -ldflags="-X github.com/andr-ii/punchy/conf.version=$(git describe --always)" github.com/andr-ii/punchy
var version string

func parseArgs() []string {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("Not enough arguments")
	}

	if args[0] == "--version" || args[0] == "-v" {
		fmt.Println(version)
		os.Exit(0)
	}

	return args
}
