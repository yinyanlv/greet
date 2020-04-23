package models

import "time"

type User struct {
	Id       uint
	Username string
	Password string
	Email    string
	Phone    uint
	CreateAt *time.Time
	UpdateAt *time.Time
}
