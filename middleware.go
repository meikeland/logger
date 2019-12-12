package logger

import (
	"bytes"
	"io"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

// LogRequest 记录gin的请求日志
func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)
		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 日志格式
		contextLogger := WithFields(Fields{
			"status": c.Writer.Status(),
			"elapse": endTime.Sub(startTime),
			"ip":     c.ClientIP(),
			"method": c.Request.Method,
			"uri":    c.Request.RequestURI,
		})

		go contextLogger.Debugf("body: %q", body)
	}
}
