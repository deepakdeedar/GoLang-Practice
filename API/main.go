package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Operator string `json:"operator"`
}

func main() {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	app.POST("/", func(c *gin.Context) {
		var json Data
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		switch json.Operator {
		case "+":
			c.JSON(http.StatusOK, gin.H{"Result": json.X + json.Y})
		case "-":
			c.JSON(http.StatusOK, gin.H{"Result": json.X - json.Y})
		case "*":
			c.JSON(http.StatusOK, gin.H{"Result": json.X * json.Y})
		case "/":
			c.JSON(http.StatusOK, gin.H{"Result": json.X / json.Y})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Please Enter a valid Operator"})
		}
	})

	app.Run(":3000")
}
