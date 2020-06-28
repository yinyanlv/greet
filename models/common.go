package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Common struct {
	gorm.Model
	CreatedBy string `gorm:"type:char(36)"`
	UpdatedBy string `gorm:"type:char(36)"`
	DeletedBy string `gorm:"type:char(36)"`
}

type CommonStrID struct {
	ID        string `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	CreatedBy string     `gorm:"type:char(36)"`
	UpdatedBy string     `gorm:"type:char(36)"`
	DeletedBy string     `gorm:"type:char(36)"`
}

type CommonTime struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
