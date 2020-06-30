package models

import (
	. "prot/db"
)

type User struct {
	CommonStrID
	Username string
	Nickname string
	Password string
	Salt     string
	Role     uint8
	Status   uint8
	Email    string
	Phone    string
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
