package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RenderEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "edit", gin.H{})
}



