package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "prot/models"
	"prot/utils"
)

func RenderRegister(c *gin.Context) {
	u := User{}
	count, err := u.Count()

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"errorCode": 599,
			"message":   err,
		})
		return
	}

	if count > 0 {
		c.Redirect(http.StatusMovedPermanently, "/404")
	}  else {
		c.HTML(http.StatusOK, "register", nil)
	}
}

func Register(c *gin.Context) {

	userReq := UserReq{}

	if err := c.BindJSON(&userReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode": 599,
			"message": err,
		})
		return
	}

	pwd, err  := utils.GenMD5Pwd(userReq.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	user := User{
		CommonStrID: CommonStrID {
			ID: utils.GenUUID(),
		},
		Username: userReq.Username,
		Nickname: userReq.Nickname,
		Password: pwd,
		Salt: utils.PwdSalt,
		Email: userReq.Email,
		Phone: userReq.Phone,
	}

	if _, err := user.Insert(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
