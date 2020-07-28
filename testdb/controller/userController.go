package controller

import "github.com/gin-gonic/gin"

func SimpleLogin(c *gin.Context) {

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": "test message",
		"data":    "hello world",
	})
}
