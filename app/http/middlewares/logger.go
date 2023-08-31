package middlewares

import (
	"bytes"
	"io"
	"time"

	"github.com/czjge/gohub/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type responseBodyWriter struct {
	// as long as it's initialized with a value implementing the interface
	gin.ResponseWriter
	// 保存响应结果
	body *bytes.Buffer
}

// 重写 ResponseWriter.Write 方法
func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// Logger 记录请求日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 response 内容
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		// 获取请求数据
		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body 是一个 buffer 对象，只能读取一次
			requestBody, _ := c.GetRawData()
			// 读取后，重新赋值 c.Request.Body ，以供后续的其他操作
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 设置开始时间
		start := time.Now()
		c.Next()

		// 开始记录日志
		cost := time.Since(start)
		responseStatus := c.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", responseStatus),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			// zap.String("time", helpers.MicrosecondsStr(cost)),
			zap.String("time", cost.String()),
		}

		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			// 请求的内容
			logFields = append(logFields, zap.String("Request Body", string(requestBody)))
			// 响应的内容
			logFields = append(logFields, zap.String("Response Body", w.body.String()))
		}

		if responseStatus > 400 && responseStatus <= 499 {
			logger.Warn("HTTP Warning "+cast.ToString(responseStatus), logFields...)
		} else if responseStatus >= 500 && responseStatus <= 599 {
			logger.Error("HTTP Error "+cast.ToString(responseStatus), logFields...)
		} else {
			logger.Debug("HTTP Access Log", logFields...)
		}
	}
}
