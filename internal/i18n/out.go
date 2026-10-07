package i18n

import (
	"fmt"
)

func Out(typeKey string, key string) {
	fmt.Println(Get(typeKey) + ": " + Get(key))
}