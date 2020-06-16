package main

import (
	"github.com/gin-gonic/gin"
	. "prot/controllers"
)

func InitRouter(r *gin.Engine) {
	r.GET("/", RenderIndex)

	r.GET("/register", RenderRegister)

	r.GET("/article", RenderArticle)
	r.POST("/article", CreateArticle)
	r.PUT("/article/:id", UpdateArticle)
	r.DELETE("/article/:id", DeleteArticle)

	r.GET("/tags", GetTags)

	r.GET("/edit/article", RenderEditArticle)

	r.NoMethod(Render404)
	r.NoRoute(Render404)
}
