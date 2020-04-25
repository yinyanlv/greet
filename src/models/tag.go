package models

import "time"

type Tag struct {
	Id string
	Text string
	CreateBy  uint
	CreateAt  *time.Time
	UpdateBy  uint
	UpdateAt  *time.Time
}