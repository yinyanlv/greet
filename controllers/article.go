package controllers

import (
	"fmt"
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
	var tags []string

	c.BindJSON(&article)
	article.Public = 1
	article.Status = 1
	article.CreatedBy = "bugong"
	article.UpdatedBy = "bugong"
	tags = c.PostFormArray("tags")

	fmt.Println(tags)
	article.Insert()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func UpdateArticle(c *gin.Context) {

}

func DeleteArticle(c *gin.Context) {

}
