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
		log.Fatalf("连接数据库失败：%v，连接字符串为：%s", err, connStr)
		return
	}

	log.Info("连接数据库成功！")

	MysqlDB.SingularTable(true)

	if etc.AppMode == "dev" {
		MysqlDB.LogMode(true)
	}
}

func Close() {
	MysqlDB.Close()
}
