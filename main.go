package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	r.NoRoute(func(ctx *gin.Context) {

		acceptString := ctx.GetHeader("Accept")
		if strings.Contains(acceptString, "text/html") {
			ctx.String(http.StatusNotFound, "404 Not Found")
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "Route not found",
			})
		}

	})

	r.Run(":8001")
}
