package main

import (
	//"fmt"

	"fmt"
	"goproject/httpd/handler"
	"io"
	"net/http"
	"os"

	"goproject/httpd/platform/files"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	file := files.New()
	// fmt.Println(file)
	// file.Add(files.Item{"hello", "fileno1"})
	// fmt.Println(file)
	r := gin.New()
	//r.Use(gin.Recovery()) //to recover from POST panic
	//r.Use(middlewares.Logger())
	r.Use(gin.Logger())
	//custom file logger

	//custom recovery
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	r.GET("/ping", handler.PingGet)
	r.GET("/files", handler.FilesGet(file))
	r.POST("/files", handler.FilesPost(file))
	//queryparsing
	r.GET("/user", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "hello %s %s", firstname, lastname)
	})
	//urlencoding
	r.POST("/post_form", func(c *gin.Context) {
		message := c.PostForm("message")
		name := c.DefaultPostForm("name", "anonymous")
		c.JSON(200, gin.H{
			"name":    name,
			"message": message,
		})
	})

	r.Run(":9000")

}
