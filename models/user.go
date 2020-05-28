package models

type User struct {
	Common
	Username string
	Nickname string
	Password string
	Role     uint8
	Status   uint8
	Email    string
	Phone    string
}

type IUser interface {
	Migrate(user *User)
}
