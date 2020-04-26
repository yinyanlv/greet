package models

import "time"

type Article struct {
	Id        string
	Title     string
	Summary   string
	Content   string
	Tags      string
	Public    bool
	Status    uint8
	ViewCount uint
	CreateBy  uint
	CreateAt  *time.Time
	UpdateBy  uint
	UpdateAt  *time.Time
}
