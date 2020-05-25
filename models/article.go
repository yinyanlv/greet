package models

import (
	"prot/db"
)

type Article struct {
	Common
	Title     string
	Summary   string
	Content   string
	Tags      string
	Public    bool
	Status    uint8
	ViewCount uint
}

func (article Article) Insert() {
	db.MysqlConn.Create(&article)
}
