package models

import . "prot/db"

type Tag struct {
	CommonTime
	ID   string
	Name string
	Sort uint8
}

type TagResp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (tag *Tag) GetTags() (tags []TagResp, err error) {
	if err = MysqlDB.Model(tag).Select([]string{"id", "name"}).Order("sort asc").Scan(&tags).Error; err != nil {
		return
	}
	return
}

func (tag *Tag) GetTagsByIDS(ids []string) (tags []Tag, err error) {
	if err = MysqlDB.Model(tag).Select([]string{"id", "name", "sort"}).Where("id in (?)", ids).Order("sort asc").Scan(&tags).Error; err != nil {
		return
	}
	return
}

func (tag *Tag) GetTagByID(id string) (res Tag, err error) {

	if err = MysqlDB.Model(tag).Where("id = (?)", id).Find(&res).Error; err != nil {
		return
	}
	return
}

func (tag *Tag) Insert() (id string, err error) {
	result := MysqlDB.Create(tag)
	if result.Error != nil {
		err = result.Error
		return
	}
	id = tag.ID
	return
}

func (tag *Tag) Update(id uint) (updatedTag TagResp, err error) {
	if err = MysqlDB.Select([]string{"id", "name"}).First(&updatedTag, id).Error; err != nil {
		return
	}

	if err = MysqlDB.Model(&updatedTag).Updates(tag).Error; err != nil {
		return
	}

	return
}

func (tag *Tag) Delete(id uint) (deletedTag TagResp, err error) {

	if err = MysqlDB.Select([]string{"id", "name"}).First(tag, id).Error; err != nil {
		return
	}

	if err = MysqlDB.Delete(tag).Error; err != nil {
		return
	}

	deletedTag.ID = tag.ID
	deletedTag.Name = tag.Name

	return
}
