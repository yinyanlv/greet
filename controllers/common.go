package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Render404(c *gin.Context) {
	c.HTML(http.StatusOK, "error", gin.H{
		"statusCode": 404,
		"message":    "未找到该页面",
	})
}
