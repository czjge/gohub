package main

import (
	"flag"
	"fmt"

	"github.com/czjge/gohub/bootstrap"
	btsConfig "github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/config"
	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	var env string
	flag.StringVar(&env, "env", "", "load .env file, eg: --env=testing loading .env.testing file")
	flag.Parse()
	config.InitConfig(env)

	router := gin.New()

	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
