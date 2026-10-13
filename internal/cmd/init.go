package cmd

import (
	"github.com/spf13/cobra"

	"Nic/internal/i18n"
	"Nic/internal/methods"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "在当前目录创建 .nic/ 仓库",
	Args:  cobra.NoArgs, // 不接受位置参数
	RunE: func(cmd *cobra.Command, args []string) error {
		// 读 --project，未指定则取当前目录名
		project, _ := cmd.Flags().GetString("project")
		if project == "" {
			project = methods.DefaultProjectName()
		}

		// 调核心层
		if err := methods.Init(project); err != nil {
			return err
		}

		i18n.Out("init_ok", ": ", project)
		i18n.Endl()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().String("project", "", "项目名（默认取当前目录名）")
}