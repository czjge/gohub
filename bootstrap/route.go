package bootstrap

import (
	"net/http"
	"strings"

	"github.com/czjge/gohub/app/http/middlewares"
	"github.com/czjge/gohub/routes"
	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {

	// global middlewares
	registerGlobalMiddleware(router)

	// register API routes
	routes.RegisterAPIRoutes(router)

	// register 404 route
	setup404Handler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		acceptString := ctx.GetHeader("Accept")
		if strings.Contains(acceptString, "text/html") {
			ctx.String(http.StatusOK, "404 Not Found")
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"error_code":    404,
				"error_message": "Route not found",
			})
		}
	})
}
