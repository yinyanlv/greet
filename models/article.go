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
	Public    uint8
	Status    uint8
	ViewCount uint
	Tags      []Tag `gorm:"many2many:article_tag;"`
}

type ArticleReq struct {
	Title   string
	Summary string
	Content string
	Tags    []string
}

func (article *Article) Insert() (id string, err error) {

	article.ID = NewV4().String()

	if err = MysqlDB.Save(article).Error; err != nil {
		return
	}

	id = article.ID

	return
}

func (article *Article) Article(id string) (res Article, err error) {
	MysqlDB.Find(&res, "id = ?", id)
	if err = MysqlDB.Model(&res).Related(&res.Tags, "Tags").Find(&res).Error; err != nil {
		return
	}
	return
}

func (article *Article) GetArticlesByPage(pageIndex uint8, pageSize uint8) (res []Article, err error) {
	if err = MysqlDB.Model(article).Offset((pageIndex - 1) * pageSize).Limit(pageIndex).Preload("Tags").Find(&res).Error; err != nil {
		return
	}

	return
}
