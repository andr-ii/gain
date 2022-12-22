package conf

import (
	"fmt"
	"os"
)

func parseArgs() []string {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("Not enough arguments")
	}

	if args[0] == "--version" || args[0] == "-v" {
		fmt.Println(CURRENT_VERSION)
		os.Exit(0)
	}

	return args
}
