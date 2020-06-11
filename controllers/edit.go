package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "prot/models"
)

func RenderEdit(c *gin.Context) {
	tag := Tag{}
	tags, err := tag.Tags()

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"statusCode": http.StatusInternalServerError,
			"message": err,
		})
		return
	}

	c.HTML(http.StatusOK, "edit", gin.H{
		"tags": tags,
	})
}
