package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"html/template"
	"net/http"
	"prot/models"
)

func RenderArticle(c *gin.Context) {
	md := []byte("## 这里是文章内容")
	content := string(markdown.ToHTML(md, nil, nil))

	c.HTML(http.StatusOK, "article", gin.H{
		"content": template.HTML(content),
	})
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
