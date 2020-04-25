package models

import "time"

type Article struct {
	Id        string
	Title     string
	SubTitle  string
	Content   string
	ViewCount uint
	Tags      []string
	CreateBy  uint
	CreateAt  *time.Time
	UpdateBy  uint
	UpdateAt  *time.Time
}
