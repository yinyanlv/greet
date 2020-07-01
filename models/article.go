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

func (article *Article) GetArticlesByTag(tag string, pageIndex uint64, pageSize uint64) (res []Article, err error) {
	if tag == "" {
		if err = MysqlDB.Model(article).
			Preload("Tags").Find(&res).
			Offset((pageIndex - 1) * pageSize).Limit(pageIndex).
			Order("created_at desc").Error;
			err != nil {
			return
		}
	} else {
		ids := make([]string, 20)
		if err = MysqlDB.Table("article_tag").Where("tag_id = ?", tag).Pluck("article_id", &ids).Error; err != nil {
			return
		}
		if len(ids) == 0 {
			return nil, nil
		}
		if err = MysqlDB.Model(article).
			Where("id in (?)", ids).
			Preload("Tags").Find(&res).
			Offset((pageIndex - 1) * pageSize).Limit(pageIndex).
			Order("created_at desc").Error;
			err != nil {
			return
		}
	}

	return
}

func (article *Article) GetArticlesByDate(date string, pageIndex uint64, pageSize uint64) (res []Article, err error) {
	if err = MysqlDB.Model(article).
		Offset((pageIndex - 1) * pageSize).Limit(pageIndex).
		Order("created_at desc").
		Preload("Tags").Find(&res).Error; err != nil {
		return
	}

	return
}
