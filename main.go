package main

import (
	"fmt"
	"os"

	"github.com/czjge/gohub/app/cmd"
	_make "github.com/czjge/gohub/app/cmd/make"
	"github.com/czjge/gohub/bootstrap"
	"github.com/czjge/gohub/pkg/console"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "Gohub",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,
		// 所有子命令都会执行以下代码
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			bootstrap.SetupConfig()
			bootstrap.SetupLogger()
			bootstrap.SetupDB()
			bootstrap.SetupRedis()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		_make.CmdMake,
		cmd.CmdMigrate,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	// cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
