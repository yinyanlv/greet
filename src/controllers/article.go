package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RenderArticle(c *gin.Context) {

	c.HTML(http.StatusOK, "article.html", gin.H{})
}