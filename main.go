package main

import (
	"fmt"

	"github.com/czjge/gohub/bootstrap"
	"github.com/czjge/gohub/config"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	bootstrap.SetupConfig()
	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.GetConfig().App.Port)
	if err != nil {
		fmt.Println(err.Error())
	}
}
