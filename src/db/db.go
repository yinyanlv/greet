package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"prot/src/models"
)

var Db *gorm.DB
var err error

func ConnDb(connStr string) {

	Db, err = gorm.Open("mysql", connStr)

	if err != nil {
		fmt.Errorf("创建数据库连接失败：%v", err)
		return
	}

	defer Db.Close()

	fmt.Println("创建数据库连接成功！")

	Db.SingularTable(true)

	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Article{})
	Db.AutoMigrate(&models.Tag{})
}
