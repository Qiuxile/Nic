package main

import (
	"os"

	"Nic/internal/args"
	"Nic/internal/data"
	"Nic/internal/i18n"
)

func main() {
	data.Init()
	
	if len(os.Args) < 2 {
		i18n.Out("error", "command_empty")
		os.Exit(1)
	}
	args.ParseArgs(os.Args)
}
