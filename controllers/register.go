package controllers

import (
	"net/http"
	. "prot/models"
	"prot/utils"

	"github.com/gin-gonic/gin"
)

func RenderRegister(c *gin.Context) {
	u := User{}
	count, err := u.Count()

	if err != nil {
		utils.HandleError(c, "html", 599, err)
		return
	}

	if count > 0 {
		c.Redirect(http.StatusMovedPermanently, "/404")
	} else {
		c.HTML(http.StatusOK, "register", nil)
	}
}

func Register(c *gin.Context) {

	userReq := UserReq{}

	if err := c.BindJSON(&userReq); err != nil {
		utils.HandleError(c, "json", 599, err)
		return
	}

	pwd, err := utils.GenMD5Pwd(userReq.Password)

	if err != nil {
		utils.HandleError(c, "json", 599, err)
		return
	}

	user := User{
		CommonStrID: CommonStrID{
			ID: utils.GenUUID(),
		},
		Username: userReq.Username,
		Nickname: userReq.Nickname,
		Password: pwd,
		Salt:     utils.PwdSalt,
		Email:    userReq.Email,
		Phone:    userReq.Phone,
	}

	if _, err := user.Insert(); err != nil {
		utils.HandleError(c, "json", 599, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
