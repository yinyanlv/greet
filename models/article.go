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
	ID      string
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

func (article *Article) Update() (id string, err error) {

	if err = MysqlDB.Table("article_tag").Where("article_id = ?", article.ID).Delete(nil).Error; err != nil {
		return
	}

	if err = MysqlDB.Save(article).Error; err != nil {
		return
	}

	id = article.ID

	return
}

func (article *Article) Delete(id string) (res string, err error) {
	res = id
	if err = MysqlDB.Table("article").Where("id = ?", id).Delete(nil).Error; err != nil {
		return
	}

	return
}

func (article *Article) Article(id string) (res Article, err error) {
	MysqlDB.Find(&res, "id = ?", id)
	if err = MysqlDB.Model(&res).Related(&res.Tags, "Tags").Find(&res).Error; err != nil {
		return
	}
	return
}

func (article *Article) AddViewCount(id string) (err error) {
	err = MysqlDB.Exec("UPDATE article SET view_count = view_count + 1 WHERE id = ?", id).Error
	return
}

func (article *Article) GetArticlesByTag(tag string, pageIndex uint64, pageSize uint64) (res []Article, err error) {
	if tag == "" {
		if err = MysqlDB.Model(article).
			Preload("Tags").
			Offset((pageIndex - 1) * pageSize).Limit(pageSize).
			Order("created_at desc").
			Find(&res).Error;
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
			Preload("Tags").
			Offset((pageIndex - 1) * pageSize).Limit(pageSize).
			Order("created_at desc").
			Find(&res).Error;
			err != nil {
			return
		}
	}

	return
}

func (article *Article) GetArticlesByKeywords(keywords string) (res []Article, err error) {
	if keywords == "" {
		return nil, nil
	}

	if err = MysqlDB.Model(article).
		Where("concat(ifnull(title, ''),ifnull(summary, ''),ifnull(content, '')) like ?", "%" + keywords + "%").
		Order("created_at desc").
		Find(&res).Error;
		err != nil {
		return
	}

	return
}

func (article *Article) GetCountByTag(tag string) (count uint64, err error) {
	if tag == "" {
		if err = MysqlDB.Model(article).
			Preload("Tags").Count(&count).Error;
			err != nil {
			return
		}
	} else {
		ids := make([]string, 20)
		if err = MysqlDB.Table("article_tag").Where("tag_id = ?", tag).Pluck("article_id", &ids).Error; err != nil {
			return
		}
		if len(ids) == 0 {
			return 0, nil
		}
		if err = MysqlDB.Model(article).
			Where("id in (?)", ids).
			Preload("Tags").Count(&count).Error;
			err != nil {
			return
		}
	}

	return
}

func (article *Article) GetArticlesByYearMonth(y string, m string, pageIndex uint64, pageSize uint64) (res []Article, err error) {
	if err = MysqlDB.Model(article).
		Where("year(created_at) = ? and month(created_at) = ?", y, m).
		Preload("Tags").Find(&res).
		Offset((pageIndex - 1) * pageSize).Limit(pageSize).
		Order("created_at desc").Error;
		err != nil {
		return
	}

	return
}

func (article *Article) GetCountByYearMonth(y string, m string) (count uint64, err error) {
	if err = MysqlDB.Model(article).
		Where("year(created_at) = ? and month(created_at) = ?", y, m).
		Preload("Tags").
		Find(&count).Error;
		err != nil {
		return
	}

	return
}
