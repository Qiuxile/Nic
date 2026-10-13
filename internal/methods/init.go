package methods

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

const (
	NicDirName   = ".nic"
	configFile   = "config.json"
	goalsDirName = "goals"
	plansDirName = "plans"
)

var ErrRepoExists = errors.New("nic repository already exists")

// 数据模型

type Config struct {
	Project string `json:"project"`
	Created string `json:"created"`
}

// Goal 一个目标，stage 内嵌
type Goal struct {
	Version int     `json:"version"`
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Status  string  `json:"status"` // todo | doing | done | cut
	Created string  `json:"created"`
	Done    *string `json:"done,omitempty"`
	Reason  string  `json:"reason,omitempty"`
	Stages  []Stage `json:"stages"`
}

type Stage struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Status  string   `json:"status"`
	Created string   `json:"created"`
	Done    *string  `json:"done,omitempty"`
	Reason  string   `json:"reason,omitempty"`
	Depends []string `json:"depends,omitempty"`
}

// Plan 一个计划，引用 goal 的 ID
type Plan struct {
	Version int      `json:"version"`
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Goals   []string `json:"goals"`
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

// Init

func Init(project string) error {
	if _, err := os.Stat(NicDirName); err == nil {
		return ErrRepoExists
	} else if !os.IsNotExist(err) {
		return err
	}

	// 一次性创建 .nic/、.nic/goals/、.nic/plans/
	for _, dir := range []string{
		NicDirName,
		filepath.Join(NicDirName, goalsDirName),
		filepath.Join(NicDirName, plansDirName),
	} {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	now := time.Now().UTC().Format(time.RFC3339)

	cfg := Config{
		Project: project,
		Created: now,
	}
	if err := writeJSON(filepath.Join(NicDirName, configFile), cfg); err != nil {
		return err
	}

	// 空目录用 .gitkeep 占位，保证 git 能跟踪
	for _, dir := range []string{goalsDirName, plansDirName} {
		p := filepath.Join(NicDirName, dir, ".gitkeep")
		if err := os.WriteFile(p, []byte{}, 0o644); err != nil {
			return err
		}
	}

	return nil
}