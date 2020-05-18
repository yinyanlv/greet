package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prot/models"
)

func RenderArticle(c *gin.Context) {

	c.HTML(http.StatusOK, "article", gin.H{})
}

func CreateArticle(c *gin.Context) {
	var article models.Article

	article.Title = c.Request.PostFormValue("title")
	article.Summary = c.Request.PostFormValue("summary")
	article.Tags = c.Request.PostFormValue("tags")
	article.Content = c.Request.PostFormValue("content")
	article.CreateBy = "bugong"

	article.Insert()

	c.JSON(http.StatusOK, gin.H{
		"code": "success",
	})
}

func UpdateArticle(c *gin.Context) {

}
