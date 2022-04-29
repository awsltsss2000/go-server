package middleware

import (
	"go-server/internal/pkg/ginx"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				url := c.Request.URL
				method := c.Request.Method
				logrus.WithFields(logrus.Fields{
					"url":    url,
					"method": method,
				}).Error(string(debug.Stack()))

				ginx.Fail(c, 500, "Internal Panic")
				// c.Abort() // 最后执行，可以不加Abort
			}
		}()
		c.Next()
	}
}
