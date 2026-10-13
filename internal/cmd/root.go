package cmd

import (
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use: "nic",
	Short: "Nic - 仓库内的项目意图账本",
	Long:  "Nic 帮你把项目想清楚，用结构化的方式记录目标、阶段和依赖。",

	// 让 main.go 用 i18n 统一输出
	SilenceErrors: true,
	SilenceUsage:  true,

	// 不带子命令时输出帮助
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}