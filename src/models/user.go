package models

import "time"

type User struct {
	Id       string
	Username string
	Nickname string
	Password string
	Role     string
	Email    string
	Phone    uint
	CreateBy string
	CreateAt *time.Time
	UpdateBy string
	UpdateAt *time.Time
}
