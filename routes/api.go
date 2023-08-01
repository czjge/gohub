package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register API routes.
func RegisterAPIRoutes(r *gin.Engine) {

	// v1 route group
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})
	}
}