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
	id := c.Param("id")
	article := models.Article{}
	a, err := article.Article(id)
	if err != nil {

		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"errorCode": 599,
			"message": err,
		})
		return
	}
	fmt.Println("%+v", a)
	fmt.Println(a.Content)
	md := []byte(a.Content)
	content := string(markdown.ToHTML(md, nil, nil))

	c.HTML(http.StatusOK, "article", gin.H{
		"title": a.Title,
		"tags": a.Tags,
		"content": template.HTML(content),
		"createdAt": a.CreatedAt,
		"updatedAt": a.UpdatedAt,
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

	fmt.Println(userId)

	var articleReq models.ArticleReq
	var tags []string

	c.BindJSON(&articleReq)
	tags = articleReq.Tags

	tag := models.Tag{}

	newTags,_ := tag.TagsByIDS(tags)

	article := models.Article{
		Title: articleReq.Title,
		Summary: articleReq.Summary,
		Tags: newTags,
		Content: articleReq.Content,
		CommonStrID: models.CommonStrID{
			CreatedBy: userId,
			UpdatedBy: userId,
		},
	}
	id, _ := article.Insert()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": id,
	})
}

func UpdateArticle(c *gin.Context) {

}

func DeleteArticle(c *gin.Context) {

}
