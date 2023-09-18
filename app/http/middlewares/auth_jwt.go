package middlewares

import (
	"github.com/czjge/gohub/app/models/user"
	"github.com/czjge/gohub/pkg/jwt"
	"github.com/czjge/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		claims, err := jwt.NewJWT().ParseToken(c)

		if err != nil {
			response.Unauthorized(c, "JWT 解析失败")
			return
		}

		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到 jwt 对应的用户")
			return
		}

		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)

		// 中间件里，当在 c.Next() 之前 return 掉，就会中断所有的后续请求
		c.Next()
	}
}
