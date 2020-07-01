package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prot/models"
	"prot/utils"
)

func RenderListByTag(c *gin.Context) {
	tag := c.Param("id")
	pageIndex, pageSize := utils.GetPage(c)

	a := models.Article{}

	articles, err := a.GetArticlesByTag(tag, pageIndex, pageSize)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"errorCode": 599,
			"message":   err,
		})
		return
	}

	count, err := a.GetCountByTag("")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"errorCode": 599,
			"articles": err,
		})
		return
	}
	pagination := utils.GetPagination(count, pageIndex, pageSize)

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
		"pagination": pagination,
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
