package models

import "time"

type Tag struct {
	Id       string
	Code     string
	Text     string
	Sort     uint32
	CreateBy string
	CreateAt *time.Time
	UpdateBy string
	UpdateAt *time.Time
}
