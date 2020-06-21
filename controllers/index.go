package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RenderIndex(c *gin.Context) {
	session := sessions.Default(c)

	userInfo := session.Get("userInfo")
	fmt.Println(userInfo)

	c.HTML(http.StatusOK, "index", gin.H{
		"userInfo": userInfo,
	})
}