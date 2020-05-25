package models

type User struct {
	Common
	Username string
	Nickname string
	Password string
	Role     string
	Status   uint8
	Email    string
	Phone    uint
}

type IUser interface {
	Migrate(user *User)
}

