package models

import (
	"github.com/jinzhu/gorm"
)

type Common struct {
	gorm.Model
	CreatedBy uint
	UpdatedBy uint
	DeletedBy uint
}
