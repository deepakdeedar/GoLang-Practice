package handler

import (
	"fmt"
	"goproject/httpd/platform/files"
	"net/http"

	"github.com/gin-gonic/gin"
)

type filepostrequest struct {
	Name string `json:"name"`
	Post string `json:"post"`
}

func FilesPost(f *files.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := filepostrequest{}
		c.Bind(&requestBody)

		//creating a panic
		if requestBody.Name == "" {
			msg, _ := fmt.Println("Panicking")
			panic(msg)
		}
		item := files.Item{
			Name: requestBody.Name,
			Post: requestBody.Post,
		}
		f.Add(item)
		c.Status(http.StatusNoContent)
	}
}
