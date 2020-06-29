package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MysqlDB *gorm.DB
var err error

func ConnDB() {
	connStr := "root:111111@tcp(127.0.0.1:3306)/prot?charset=utf8&parseTime=True&loc=Local"

	MysqlDB, err = gorm.Open("mysql", connStr)

	if err != nil {
		fmt.Errorf("创建数据库连接失败：%v", err)
		return
	}

	fmt.Println("创建数据库连接成功！")

	MysqlDB.SingularTable(true)
	MysqlDB.LogMode(true)
}

func Close()  {
	MysqlDB.Close()
}
