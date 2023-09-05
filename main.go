package main

import (
	"fmt"

	"github.com/czjge/gohub/bootstrap"
	"github.com/czjge/gohub/config"

	"github.com/gin-gonic/gin"
)

func main() {

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	bootstrap.SetupConfig()
	bootstrap.SetupLogger()
	bootstrap.SetupDB()
	bootstrap.SetupRedis()
	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.GetConfig().App.Port)
	if err != nil {
		fmt.Println(err.Error())
	}
}
