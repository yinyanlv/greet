package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RenderIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "views/index", gin.H{})
}