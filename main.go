package main

import (
	"github.com/gin-gonic/gin"
	"prot/models"

	. "prot/db"
	. "prot/utils"
)

func main() {
	ConnDb()
	MysqlDb.AutoMigrate(&models.User{})
	MysqlDb.AutoMigrate(&models.Article{})
	MysqlDb.AutoMigrate(&models.Tag{})

	r := gin.Default()
	commonTpls := GetFileList("./templates/common", ".html")

	r.Static("/public", "./public")
	r.HTMLRender = LoadTemplateFiles("./templates", ".html", commonTpls)
	InitRouter(r)

	r.Run(":9000")
}
