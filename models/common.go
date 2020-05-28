package models

import "time"

type Common struct {
	ID        uint
	CreatedBy string
	CreateAt  time.Time
	UpdateBy  string
	UpdateAt  time.Time
	DeleteBy  string
	DeleteAt  time.Time
}
