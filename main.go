package main

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	. "prot/db"
	"prot/middlewares"
	"prot/models"
	. "prot/utils"
	"time"
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

	r.SetFuncMap(template.FuncMap{
		"formatDatetime": func(dt time.Time) string {
			return dt.Format("2020-01-01 01:01:01")
		},
		"year": func(dt time.Time) int {
			return dt.Year()
		},
		"month": func(dt time.Time) int {
			return int(dt.Month())
		},
	})

	InitRouter(r)

	r.Run(":9000")
}
