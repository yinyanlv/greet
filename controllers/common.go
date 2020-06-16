package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prot/utils"
)

func Render404(c *gin.Context) {
	if utils.IsAjax(c) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "未找到该页面",
		})
	} else {
		c.HTML(http.StatusNotFound, "error", gin.H{
			"message": "未找到该页面",
		})
	}
}
