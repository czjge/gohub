package jwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/app"
	"github.com/czjge/gohub/pkg/logger"
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt/v5"
)

var (
	ErrHeaderEmpty            error = errors.New("token is required")
	ErrTokenExpiredMaxRefresh error = errors.New("token has expired max refresh time")
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

	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// 刷新 token
func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {

	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	token, err := jwt.parseTokenString(tokenString)

	// 如果 error 是 jwtpkg.ErrTokenExpired，则可以刷新 token
	if err != nil && err != fmt.Errorf("%w: %w", jwtpkg.ErrTokenInvalidClaims, jwtpkg.ErrTokenExpired) {
		return "", err
	}

	claims := token.Claims.(*JWTCustomClaims)

	// 检查是否超过最大可刷新时间
	x := app.TimenowInTimezone().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt.Time.Unix() > x {
		claims.RegisteredClaims.ExpiresAt = jwtpkg.NewNumericDate(jwt.expireAtTime())
		return jwt.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
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
		return "", jwtpkg.ErrTokenMalformed
	}

	return parts[1], nil
}

func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwtpkg.Token) (any, error) {
		return jwt.SignKey, nil
	})
}
