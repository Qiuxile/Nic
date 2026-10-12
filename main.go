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
		i18n.Endl()
		os.Exit(1)
	}

	if err := args.ParseArgs(os.Args); err != nil {
		// 解析参数
		i18n.Out("error", err.Error())
		i18n.Endl()
		os.Exit(1)
	}
}
