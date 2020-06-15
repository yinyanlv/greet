package models

type User struct {
	CommonStrID
	Username string
	Nickname string
	Password string
	Role     uint8
	Status   uint8
	Email    string
	Phone    string
}
