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
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}

type CommonTime struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
