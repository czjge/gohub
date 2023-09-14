package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/app"
	"github.com/czjge/gohub/pkg/logger"
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalFormed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问")
	ErrHeaderMalFormed        error = errors.New("请求头中 Authorization 格式有误")
)

// jwt 对象
type JWT struct {
	// 加密 jwt 密匙
	SignKey []byte
	// 刷新 token 最大过期时间
	MaxRefresh time.Duration
}

// jwt 自定义载荷
type JWTCustomClaims struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`

	// jwt 基础 struct
	jwtpkg.RegisteredClaims
}

func NewJWT() *JWT {
	config := config.GetConfig().Jwt
	return &JWT{
		SignKey:    []byte(config.SignKey),
		MaxRefresh: time.Duration(config.MaxRefreshTime) * time.Minute,
	}
}

// 生成 token
func (jwt *JWT) IssueToken(userID string, userName string) string {

	expireAtTime := jwt.expireAtTime()
	claims := JWTCustomClaims{
		userID,
		userName,
		expireAtTime.Unix(),
		jwtpkg.RegisteredClaims{
			NotBefore: jwtpkg.NewNumericDate(app.TimenowInTimezone()),
			IssuedAt:  jwtpkg.NewNumericDate(app.TimenowInTimezone()), // 后续刷新 token 不会更新
			ExpiresAt: jwtpkg.NewNumericDate(expireAtTime),
			Issuer:    config.GetConfig().App.Name,
		},
	}

	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}

	return token
}

// 解析 token
func (jwt *JWT) ParseToken(c *gin.Context) (*JWTCustomClaims, error) {

	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	token, err := jwt.parseTokenString(tokenString)

	if err != nil {

		// validationErr, ok := err.(jwtpkg.)
	}
}

// 创建 token（内部调用）
func (jwt *JWT) createToken(claims JWTCustomClaims) (string, error) {
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

func (jwt *JWT) expireAtTime() time.Time {

	timenow := app.TimenowInTimezone()

	config := config.GetConfig()

	var expireTime int64
	if config.App.Debug {
		expireTime = int64(config.Jwt.DebugExpireTime)
	} else {
		expireTime = int64(config.Jwt.ExpireTime)
	}

	expire := time.Duration(expireTime) * time.Minute
	return timenow.Add(expire)
}

// Authorization:Bearer xxxxx
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {

	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalFormed
	}

	return parts[1], nil
}

func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwtpkg.Token) (any, error) {
		return jwt.SignKey, nil
	})
}
