package cmd

import (
	"github.com/czjge/gohub/bootstrap"
	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/console"
	"github.com/czjge/gohub/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// 运行 Web 服务命令
var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.GetConfig().App.Port)
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}
}
