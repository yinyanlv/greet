package models

import (
	. "prot/db"
)

type User struct {
	CommonStrID
	Username string `gorm:"type:varchar(20) not null"`
	Nickname string `gorm:"type:varchar(20)"`
	Password string `gorm:"type:varchar(40) not null"`
	Salt     string `gorm:"type:varchar(20) not null"`
	Role     uint8  `gorm:"type:tinyint unsigned"`
	Status   uint8  `gorm:"type:tinyint unsigned"`
	Email    string `gorm:"type:varchar(30)"`
	Phone    string `gorm:"type:varchar(20)"`
}

type UserReq struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (user User) Insert() (id string, err error) {
	result := MysqlDB.Create(&user)

	if err = result.Error; err != nil {
		return
	}

	id = user.ID

	return
}

func (user User) IsExist() (bool, User) {
	var count uint8
	var u User

	if MysqlDB.Model(&user).Select([]string{
		"id",
		"username",
		"nickname",
		"role",
		"status",
		"email",
		"phone",
	}).Where(&user).First(&u).Count(&count).Error != nil {
		return false, User{}
	}

	if count == 1 {
		return true, u
	} else {
		return false, User{}
	}
}
