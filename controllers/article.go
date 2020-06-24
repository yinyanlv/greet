package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
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

func RenderEditArticle(c *gin.Context) {
	tag := models.Tag{}
	tags, err := tag.Tags()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"message": err,
		})
		return
	}

	id := c.Query("id")

	if id == "" {
		c.HTML(http.StatusOK, "edit", gin.H{
			"tags": tags,
		})
	} else {

	}
}

func CreateArticle(c *gin.Context) {
	session := sessions.Default(c)
	userInfo := session.Get("userInfo").(models.User)
	userId := userInfo.ID

	var article models.Article
	var tags []string

	c.BindJSON(&article)
	article.Public = 1
	article.Status = 1
	article.CreatedBy = userId
	article.UpdatedBy = userId
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
