package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prot/models"
)

func GetTags(c *gin.Context)  {
	tag := models.Tag{}
	tags, err := tag.Tags()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"code": "db-error",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": tags,
	})
}

