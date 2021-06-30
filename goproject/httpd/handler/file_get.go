package handler

import (
	"goproject/httpd/platform/files"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FilesGet(f *files.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := f.GetAll()
		c.JSON(http.StatusOK, result)
	}
}
