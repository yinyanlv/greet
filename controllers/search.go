package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"prot/models"
	"prot/utils"
)

func RenderSearch(c *gin.Context) {
	keywords := c.Query("k")
	a := models.Article{}

	articles, err := a.GetArticlesByKeywords(keywords)
	if err != nil {
		utils.HandleError(c, "html", 599, err)
		return
	}

	session := sessions.Default(c)
	userInfo := session.Get("userInfo")

	c.HTML(http.StatusOK, "search", gin.H{
		"pageCode": "search",
		"userInfo": userInfo,
		"articles": articles,
		"keywords": keywords,
	})
}
