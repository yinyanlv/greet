package models

import (
	"prot/db"
	"time"
)

type Article struct {
	Id        string
	Title     string
	Summary   string
	Content   string
	Tags      string
	Public    bool
	Status    uint8
	ViewCount uint
	CreateBy  string
	CreateAt  *time.Time
	UpdateBy  string
	UpdateAt  *time.Time
}

func (article Article) Insert() {
	db.Db.Create(&article)
}
