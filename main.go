package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	. "prot/db"
	. "prot/migrations"
	. "prot/utils"
)

func main() {
	ConnDB()

	InitDB()

	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))

	r.Static("/public", "./public")

	commonTpls := GetFileList("./templates/common", ".html")
	r.HTMLRender = LoadTemplateFiles("./templates", ".html", commonTpls)
	InitRouter(r)

	r.Run(":9000")
}
