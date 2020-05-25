package models

import "time"

type Common struct {
	ID string
	CreateBy  string
	CreateAt  *time.Time
	UpdateBy  string
	UpdateAt  *time.Time
}
