package config

import (
	"os"
	"errors"
	"path/filepath"
	"encoding/json"
)

var configPath string

type Config struct {
	Language string `json:"language"`
}

func init() {
	// 初始化config文件路径
	exe, err := os.Executable()
	if err != nil { panic(err) }

	if resolve, err := filepath.EvalSymlinks(exe); err == nil {
		exe = resolve
	}

	configPath = filepath.Join(filepath.Dir(exe), "config.json")
}

func defaultConfig() Config {
	// 默认配置文件信息
	return Config{
		Language: "en",
	}
}

func InitConfigFile() error {
	// 创建/修改配置文件
	data, err := json.MarshalIndent(defaultConfig(), "", "  ")
	if err != nil { return err }

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return err
	}
	return nil
}

func ReadConfigFile() (Config, error) {
	// 读取config
	fileContent, err := os.ReadFile(configPath)
	if err != nil { return defaultConfig(), err }

	// json解析config
	var cfg Config
	if err := json.Unmarshal(fileContent, &cfg); err != nil { 
		return defaultConfig(), err
	}

	return cfg, nil
} 

func LoadConfigFile() (Config, error) {
	// 检查配置文件
	_, err := os.Stat(configPath)

	if errors.Is(err, os.ErrNotExist) {
		if err := InitConfigFile(); err != nil {
			return defaultConfig(), err
		}
	}
	return ReadConfigFile()
}