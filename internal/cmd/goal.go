package cmd

import (
	"errors"

	"github.com/spf13/cobra"

	"Nic/internal/i18n"
	"Nic/internal/methods"
)

var goalCmd = &cobra.Command{
	Use: "goal",
	Short: "当前目录下 .nic 仓库的 goal 操作",
}

var goalAddCmd = &cobra.Command{
	Use:   "add",
	Short: "在当前目录下的.nic仓库中添加一个goal",
	Args:  cobra.NoArgs, // 不接受位置参数
	RunE: func(cmd *cobra.Command, _ []string) error {
		// 读取title
		title, err := cmd.Flags().GetString("title")
		if err != nil {
			return err
		}
		if title == "" {
			return errors.New(i18n.Get("content_empty"))
		}

		// 读取
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		if id == "" {
			return errors.New(i18n.Get("content_empty"))
		}

		// 调核心层
		if err := methods.AddGoal(title, id); err != nil {
			return err
		}
		return nil
	},
}

var goalListCmd = &cobra.Command{
	Use: "list",
	Short: "列出当前目录下的.nic仓库中所有goal",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return methods.ListGoal()
	},
}


func init() {
	rootCmd.AddCommand(goalCmd)
	goalCmd.AddCommand(goalAddCmd, goalListCmd)

	goalAddCmd.Flags().String("title", "", "目标标题")
	goalAddCmd.Flags().String("id", "", "目标id")
}