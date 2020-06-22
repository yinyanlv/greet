package middlewares

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"prot/utils"
)

var checkLoginUrls = []string{
	"/edit/article",
}

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		url := c.Request.URL.String()
		fmt.Println(url)
		session := sessions.Default(c)
		if u := session.Get("userInfo"); u != nil {
			c.Next()
			return
		}

		if isNeedLogin(url) {
			if utils.IsAjax(c) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "用户未登录",
				})
			} else {
				c.HTML(http.StatusUnauthorized, "error", gin.H{
					"errorCode": "401",
					"message": "用户未登录",
				})
			}
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func isNeedLogin(url string) bool {
	for _, item := range checkLoginUrls {
		if item == url {
			return true
		}
	}
	return false
}
