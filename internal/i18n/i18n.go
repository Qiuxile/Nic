package i18n

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"Nic/internal/config"
)

var langlist map[string]string

func must(err error) {
	if err != nil { panic(err) }
}

func init() {
	exe, err := os.Executable()
	must(err)

	if resolve, err := filepath.EvalSymlinks(exe); err == nil {
		exe = resolve
	}

	targetDir := filepath.Join(filepath.Dir(exe), "i18n")
	entries, err := os.ReadDir(targetDir)
	if err != nil {
		log.Println("Read i18n dir error: "+err.Error())
		return
	}
	cfg, err := config.LoadConfigFile()
	if err != nil {
		log.Println("Load config file error: "+err.Error())
		return
	}
	target := cfg.Language + ".json"
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if entry.Name() == target {
			if err := LoadLangList(filepath.Join(targetDir, entry.Name())); err != nil {
				log.Println("Load langlist file error: " + err.Error())
				return
			}
			break
		}
	}
}

func LoadLangList(targetPath string) error {
	content, err := os.ReadFile(targetPath)
	if err != nil { return err }

	if err := json.Unmarshal(content, &langlist); err != nil {
		return err
	}
	return nil
}

func Get(key string) string {
	if value, ok := langlist[key]; ok {
		return value
	}
	// log.Printf("Not found value of %q in langlist", key)
	return key
}