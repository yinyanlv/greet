package main

import (
	"github.com/gin-gonic/gin"
	"prot/src/controllers"
)

func InitControllers(r *gin.Engine) {
	r.GET("/", controllers.RenderIndex)
	r.GET("/article", controllers.RenderArticle)
}
