package methods

import (
	"os"
	"path/filepath"
	"encoding/json"
)

const (
	NicDirName   = ".nic"
	configFile   = "config.json"
	goalsDirName = "goals"
	plansDirName = "plans"
)

// 数据模型
type Config struct {
	Project string `json:"project"`
	Created string `json:"created"`
}

// Goal 一个目标，stage 内嵌
type Goal struct {
	ID 		string	 `json:"id"`
	Title	string	 `json:"title"`
	Descripts string `json:"descripts"`
	Stages	[]Stage	 `json:"stages"`
	Status	string	 `json:"status"`
	Created	string	 `json:"created"`
}

type Stage struct {
	ID 		string	 `json:"id"`
	Title	string	 `json:"title"`
	Descripts string `json:"descripts"`
	Stages	[]Stage	 `json:"stages"`
	Status	string	 `json:"status"`
	Created	string	 `json:"created"`
}

// Plan 一个计划，引用 goal 的 ID
type Plan struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Descripts string `json:"descripts"`
	Goals   []string `json:"goals"`
	Status	string	 `json:"status"`
	Created string   `json:"created"`
}


// 工具函数
func DefaultProjectName() string {
	wd, err := os.Getwd()
	if err != nil {
		return "unnamed"
	}
	return filepath.Base(wd)
}

func writeJSON(path string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0o644)
}

func CherkDir(DirPath string) error {
	// 检查文件夹是否存在
	if _, err := os.Stat(DirPath); err == nil {
		return err
	} else if !os.IsNotExist(err) {
		return err
	}
	return nil
}

func CreateDir(path []string) error {
	for _, dir := range path {
		if err := os.Mkdir(dir, 0o755); err != nil {
			return err
		}
	}
	return nil
}