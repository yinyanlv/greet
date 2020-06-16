package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RenderRegister(c *gin.Context) {

	c.HTML(http.StatusOK, "register", gin.H{
	})
}
