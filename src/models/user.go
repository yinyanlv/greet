package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	. "prot/src/db"
)

type User struct {
	Id       string
	Username string
	Nickname string
	Password string
	Role     string
	Status   uint8
	Email    string
	Phone    uint
	CreateBy string
	CreateAt *time.Time
	UpdateBy string
	UpdateAt *time.Time
}

func GetUsers(c *gin.Context) {
	var users []User
	if err := Db.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}
}
