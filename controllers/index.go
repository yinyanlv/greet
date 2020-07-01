package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"prot/models"
	"strconv"
)

func RenderIndex(c *gin.Context) {
	session := sessions.Default(c)
	pageIndex, _ := strconv.ParseUint(c.Query("page"), 10, 64)
	pageSize, _ := strconv.ParseUint(c.Query("size"), 10, 64)
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}
	a := models.Article{}
	var articles []models.Article
	var err error

	articles, err = a.GetArticlesByTag("", pageIndex, pageSize)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"errorCode": 599,
			"articles": err,
		})
		return
	}

	userInfo := session.Get("userInfo")

	c.HTML(http.StatusOK, "index", gin.H{
		"pageCode": "index",
		"userInfo": userInfo,
		"articles": articles,
	})
}