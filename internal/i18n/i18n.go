package i18n

import (
	"os"
	_ "encoding/json"
)

var (
	Language string
)


func getSystemLanguage() string {
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = "en"
	}
	return lang
}

func LoadI18nDir(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil { return err }
	
	for _, entry := range entries {
		if entry.IsDir() { continue }
	}
	
	return nil
}

func Get(key string) string {
	return key
}