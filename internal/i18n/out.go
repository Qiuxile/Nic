package i18n

import (
	"fmt"
)

func Out(keys ...string) error {
	for _, key := range keys {
		if _, err := fmt.Print(Get(key)); err != nil {
			Out("error", ": ", err.Error())
		}
	}
	return nil
}

func Endl() {
	fmt.Print("\n\n")
}