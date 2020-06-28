package models

import (
	. "github.com/satori/go.uuid"
	. "prot/db"
)

type Article struct {
	CommonStrID
	Title     string `gorm:"type:varchar(200) not null"`
	Summary   string `gorm:"type:varchar(200)"`
	Content   string `gorm:"type:text not null"`
	Public    uint8  `gorm:"type:tinyint unsigned"`
	Status    uint8  `gorm:"type:tinyint unsigned"`
	ViewCount uint   `gorm:"type:int unsigned"`
	Tags      []Tag  `gorm:"many2many:article_tag"`
}

func (article *Article) Insert() (id string, err error) {

	article.ID = NewV4().String()

	if err = MysqlDB.Create(article).Error; err != nil {
		return
	}

	id = article.ID

	return
}
