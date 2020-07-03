package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	. "prot/models"
	"prot/utils"
)

func RenderLogin(c *gin.Context) {

	c.HTML(http.StatusOK, "login", nil)
}

func Login(c *gin.Context) {
	session := sessions.Default(c)

	userReq := UserReq{}

	if err := c.BindJSON(&userReq); err != nil {
		utils.HandleError(c, "json", 599, err)
		return
	}

	pwd, err  := utils.GenMD5Pwd(userReq.Password)

	if err != nil {
		utils.HandleError(c, "json", 599, err)
		return
	}

	user := User{
		Username: userReq.Username,
		Password: pwd,
	}

	if isExist, u := user.IsExist(); isExist {
		session.Set("userInfo", u)
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户名或密码不正确",
		})
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/")
}
