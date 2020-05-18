package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"prot/models"

	. "prot/db"
	. "prot/utils"
)

func main() {
	ConnDb()
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Article{})
	Db.AutoMigrate(&models.Tag{})

	r := gin.Default()
	commonTpls := GetFileList("./templates/common", ".html")

	r.Static("/public", "./public")
	r.HTMLRender = LoadTemplateFiles("./templates", ".html", commonTpls)
	InitRouter(r)

	r.Run(":9000")
}
