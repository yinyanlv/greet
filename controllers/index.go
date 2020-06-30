package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"prot/models"
)

func RenderIndex(c *gin.Context) {
	session := sessions.Default(c)
	a := models.Article{}
	articles, err := a.GetArticlesByPage(1, 20)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"errorCode": 599,
			"articles": err,
		})
		return
	}

	userInfo := session.Get("userInfo")
	fmt.Println(userInfo)
	fmt.Println(articles)

	c.HTML(http.StatusOK, "index", gin.H{
		"userInfo": userInfo,
		"articles": articles,
	})
}