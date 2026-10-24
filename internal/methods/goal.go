package methods

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"Nic/internal/i18n"
)

func CherkGoals() error {
	// 检查.nic是否存在
	if err := CherkDir(NicDirName); err != nil {
		return err
	}
	// 检查.nic/goals是否存在
	if err := CherkDir(filepath.Join(NicDirName, goalsDirName)); err != nil {
		return err
	}
	return nil
}

func ListGoal() error {
	if err := CherkGoals(); err != nil {
		return err
	}

	if entries, err := os.ReadDir(filepath.Join(NicDirName, goalsDirName)); err != nil {
		return err
	} else {
		for _, entry := range entries {
			fmt.Println(entry)
		}
	}

	return nil
}

func AddGoal(title string, id string) error {
	if err := CherkGoals(); err != nil {
		return nil
	}

	DefaultGoal := Goal{
		ID: id,
		Title: title,
		Created: time.Now().UTC().Format(time.RFC3339),
		Status: "todo",
	}

	if err := writeJSON(filepath.Join(NicDirName, goalsDirName, id+".json"), DefaultGoal); err != nil {
		return err
	}

	i18n.Out("goal_ok", ": ", title)
	i18n.Endl()

	return nil
}