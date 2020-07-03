package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"prot/etc"
)

var MysqlDB *gorm.DB
var err error

func ConnDB() {
	c := etc.AppConfig
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.MySQL.Username,
		c.MySQL.Password,
		c.MySQL.Host,
		c.MySQL.Port,
		c.MySQL.DB,
	)

	MysqlDB, err = gorm.Open("mysql", connStr)

	if err != nil {
		log.Fatalf("创建数据库连接失败：%v", err)
		return
	}

	log.Info("创建数据库连接成功！")

	MysqlDB.SingularTable(true)

	if etc.AppMode == "dev" {
		MysqlDB.LogMode(true)
	}
}

func Close() {
	MysqlDB.Close()
}
