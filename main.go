package main

import (
	"github.com/gin-gonic/gin"
	. "prot/db"
	. "prot/utils"
)

func main() {
	ConnDb()

	r := gin.Default()
	commonTpls := GetFileList("./templates/common", ".html")

	r.Static("/public", "./public")
	r.HTMLRender = LoadTemplateFiles("./templates", ".html", commonTpls)
	InitRouter(r)

	r.Run(":9000")
}
