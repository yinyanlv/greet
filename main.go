package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	. "prot/db"
	. "prot/utils"
)

func main() {

	connStr := "root:111111@tcp(127.0.0.1:3306)/prot?charset=utf8&parseTime=True&loc=Local"
	ConnDb(connStr)
	Migrate()

	r := gin.Default()
	commonTpls := GetFileList("./templates/common", ".html")

	r.Static("/public", "./public")
	r.HTMLRender = LoadTemplateFiles("./templates", ".html", commonTpls)
	InitRouter(r)

	r.Run(":9000")
}
