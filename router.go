package main

import (
	"github.com/gin-gonic/gin"
	. "prot/controllers"
)

func InitRouter(r *gin.Engine) {
	r.GET("/", RenderIndex)

	r.GET("/article", RenderArticle)
	r.POST("/article", CreateArticle)
	r.PUT("/article", UpdateArticle)
	r.GET("/edit", RenderEdit)
}
