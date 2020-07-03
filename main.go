package main

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	. "prot/db"
	"prot/middlewares"
	"prot/models"
	. "prot/utils"
)

func main() {
	ConnDB()

	gob.Register(models.User{})

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("yyl", store))

	r.Use(middlewares.Auth())

	r.Static("/public", "./public")

	commonTpls := GetFileList("./templates/common", ".html")
	r.HTMLRender = LoadTemplateFiles("./templates", ".html", commonTpls)

	InitRouter(r)

	r.Run(":9000")
}
