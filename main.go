package main

import (
	"os"

	"Nic/internal/args"
	"Nic/internal/i18n"
)

const (
	version string = "0.0.1"
	Author string = "Surile"
)


func main() {
	if len(os.Args) < 2 {
		// 空处理
		i18n.Out("error", ": ", "command_empty")
		os.Exit(1)
	}

	err := args.ParseArgs(os.Args)
	if err != nil {
		i18n.Out("error", err.Error())
		os.Exit(1)
	}
}
