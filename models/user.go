package models

type User struct {
	CommonStrID
	Username string `gorm:"type:varchar(20) not null"`
	Nickname string `gorm:"type:varchar(20)"`
	Password string `gorm:"type:varchar(40) not null"`
	Role     uint8  `gorm:"type:tinyint unsigned"`
	Status   uint8  `gorm:"type:tinyint unsigned"`
	Email    string `gorm:"type:varchar(30)"`
	Phone    string `gorm:"type:varchar(20)"`
}
