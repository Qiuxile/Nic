package args

import (
	"fmt"

	"Nic/internal/i18n"
	_"Nic/internal/methods"
)

func outArgs(args []string) error {
	for i, arg := range args {
		if err := i18n.Out("参数[", fmt.Sprint(i), "]为: ", string(arg), "\n"); err != nil {
			return err
		}
	}
	return nil
}

func ParseArgs(args []string) error {
	// 在此处实现参数解析逻辑

	if args[1] == "version" {
		i18n.Out("version", ": ", "v0.0.1")
	} else {
		i18n.Out("error", ": ", "dont have", ` "`, args[1], `" `,  "command", "!")
	}
	i18n.Endl()

	return nil
}