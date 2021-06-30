package middlewares

import (
	"github.com/gin-gonic/gin"
)

func ErrorLogger() gin.HandlerFunc {

	return gin.ErrorLoggerT(1 << 1)
}
