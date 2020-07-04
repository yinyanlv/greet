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
	var hostStr string

	if c.MySQL.Port != 0 {
		hostStr = fmt.Sprintf("%s:%d", c.MySQL.Host, c.MySQL.Port)
	} else {
		hostStr = c.MySQL.Host
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.MySQL.Username,
		c.MySQL.Password,
		hostStr,
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
