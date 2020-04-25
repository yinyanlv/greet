package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	. "prot/src/db"
)

func main() {

	connStr := "root:111111@tcp(127.0.0.1:3306)/prot?charset=utf8&parseTime=True&loc=Local"
	ConnDb(connStr)

	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("views/*")
	InitControllers(r)

	r.Run(":1124")
}
