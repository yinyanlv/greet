package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	. "prot/src/models"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "root:111111@tcp(127.0.0.1:3306)/prot?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Errorf("创建数据库连接失败：%v", err)
	}

	defer db.Close()

	fmt.Println("创建数据库连接成功！")

	db.SingularTable(true)

	db.AutoMigrate(&User{})

	db.Create(&User{
		Username: "root",
		Password: "111111",
		Email: "abc@d.com",
	})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!",
		})
	})

	r.GET("/users", GetUsers)

	r.Run()
}

func GetUsers (c *gin.Context) {
	var users []User
	if err := db.Find(&users).Error; err!= nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}


}
