package middleware

import (
	"bytes"
	"github.com/GoldenLeeK/blog-service/global"
	"github.com/GoldenLeeK/blog-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}

		global.Logger.WithFields(fields).Infof("access log:method:%s,status_code:%d,begin_time:%d,end_time:%d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime)

	}
}
