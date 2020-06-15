package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Common struct {
	gorm.Model
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}

type CommonStrID struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	CreatedBy uint
	UpdatedBy uint
	DeletedBy uint
}
