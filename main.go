package main

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	. "prot/db"
	"prot/etc"
	"prot/middlewares"
	"prot/models"
	. "prot/utils"
	"strconv"
)

func main() {
	ConnDB()
	gob.Register(models.User{})

	if etc.AppMode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("yyl", store))

	r.Use(middlewares.Auth())

	r.Static("/public", "./public")

	commonTpls := GetFileList("./templates/common", ".html")
	r.HTMLRender = LoadTemplateFiles("./templates", ".html", commonTpls)

	InitRouter(r)

	r.Run(":" + strconv.Itoa(etc.AppConfig.Port))
}
