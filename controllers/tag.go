package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prot/models"
	"strconv"
)

func RenderListByTag(c *gin.Context) {
	tag := c.Param("id")
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

	articles, err = a.GetArticlesByTag(tag, pageIndex, pageSize)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"errorCode": 599,
			"message":   err,
		})
		return
	}

	t := models.Tag{}
	tagObj, err := t.GetTagByID(tag)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"errorCode": 599,
			"message":   err,
		})
		return
	}

	c.HTML(http.StatusOK, "list", gin.H{
		"pageCode":  tag,
		"pageTitle": tagObj.Name,
		"articles":  articles,
	})
}

func GetTags(c *gin.Context) {
	tag := models.Tag{}
	tags, err := tag.GetTags()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    tags,
	})
}
