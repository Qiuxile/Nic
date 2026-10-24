package methods

import (
	"os"
	"path/filepath"
	"time"
)

func CreateDotNic() error {
	// 一次性创建 .nic/、.nic/goals/、.nic/plans/
	paths := []string{
		NicDirName,
		filepath.Join(NicDirName, goalsDirName),
		filepath.Join(NicDirName, plansDirName),
	}

	if err := CreateDir(paths); err != nil {
		return err
	}
	return nil
}


func WriteDefaultConfig(project string) error {
	// 写入默认Config文件信息
	now := time.Now().UTC().Format(time.RFC3339)
	cfg := Config{
		Project: project,
		Created: now,
	}
	if err := writeJSON(filepath.Join(NicDirName, configFile), cfg); err != nil {
		return err
	}
	return nil
}

func CreateGitKeep() error {
	// 空目录用 .gitkeep 占位，保证 git 能跟踪
	for _, dir := range []string{goalsDirName, plansDirName} {
		p := filepath.Join(NicDirName, dir, ".gitkeep")
		if err := os.WriteFile(p, []byte{}, 0o644); err != nil {
			return err
		}
	}
	return nil
}


func Init(project string) error {
	if err := CherkDir(NicDirName); err != nil {
		return err
	}
	if err := CreateDotNic(); err != nil {
		return err
	}
	if err := WriteDefaultConfig(project); err != nil {
		return err
	}
	if err := CreateGitKeep(); err != nil {
		return err
	}

	return nil
}