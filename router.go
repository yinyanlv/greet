package main

import (
	"github.com/gin-gonic/gin"
	. "prot/controllers"
)

func InitRouter(r *gin.Engine) {
	r.GET("/", RenderIndex)

	r.GET("/register", RenderRegister)
	r.POST("/register", Register)

	r.GET("/login", RenderLogin)
	r.POST("/login", Login)

	r.GET("/logout", Logout)

	r.GET("/article/:id", RenderArticle)
	r.POST("/article", CreateArticle)
	r.PUT("/article", UpdateArticle)
	r.DELETE("/article/:id", DeleteArticle)

	r.GET("/tags", GetTags)
	r.GET("/tag/:id", RenderListByTag)

	r.GET("/archive", RenderArchive)
	r.GET("/archive/:year/:month", RenderArchiveByYearMonth)

	r.GET("/edit-article", RenderEditArticle)

	r.NoMethod(Render404)
	r.NoRoute(Render404)
}
