package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PingGet(c *gin.Context) {
	timer1 := time.NewTimer(5 * time.Second)
	<-timer1.C
	c.JSON(http.StatusOK, gin.H{
		"Hello": " found me",
	})
}
