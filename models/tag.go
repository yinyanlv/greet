package models

import . "prot/db"

type Tag struct {
	Common
	Code string
	Name string
	Sort uint32
}

type TagResult struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (tag *Tag) Tags() (tags []TagResult, err error) {

	if err = MysqlDB.Model(tag).Select([]string{"code", "name"}).Scan(&tags).Error; err != nil {
		return
	}
	return
}

func (tag Tag) Insert() (id uint, err error) {
	result := MysqlDB.Create(&tag)
	if result.Error != nil {
		err = result.Error
		return
	}
	id = tag.ID
	return
}

func (tag *Tag) Update(id uint) (updatedTag TagResult, err error) {
	if err = MysqlDB.Select([]string{"code", "name"}).First(&updatedTag, id).Error; err != nil {
		return
	}

	if err = MysqlDB.Model(&updatedTag).Updates(tag).Error; err != nil {
		return
	}

	return
}

func (tag *Tag) Delete(id uint) (deleledTag TagResult, err error) {

	if err = MysqlDB.Select([]string {"code", "name"}).First(tag, id).Error; err != nil {
		return
	}

	if err = MysqlDB.Delete(tag).Error; err != nil {
		return
	}

	deleledTag.Code = (*tag).Code
	deleledTag.Name = (*tag).Name

	return
}
