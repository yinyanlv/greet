package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "prot/models"
	"prot/utils"
)

func RenderRegister(c *gin.Context) {

	c.HTML(http.StatusOK, "register", gin.H{
	})
}

func Register(c *gin.Context) {

	userReq := UserReq{}

	if err := c.BindJSON(&userReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	pwd, err  := utils.GenMD5Pwd(userReq.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
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
