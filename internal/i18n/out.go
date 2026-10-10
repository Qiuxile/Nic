package i18n

import (
	"fmt"
)

func Out(keys ...string) {
	for _, key := range keys {
		fmt.Print(Get(key))
	}
}
