package models

import (
	. "github.com/satori/go.uuid"
	. "prot/db"
)

type Article struct {
	CommonStrID
	Title     string
	Summary   string
	Content   string
	Public    bool
	Status    uint8
	ViewCount uint
}

func (article *Article) Insert() (id string, err error) {

	article.ID = NewV4().String()

	if err = MysqlDB.Create(article).Error; err != nil {
		return
	}

	id = article.ID

	return
}
