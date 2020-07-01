package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"prot/utils"
	"regexp"
	"strings"
)

func init() {
	for i, item := range checkLoginUrls {
		if strings.Index(item, "/:") > -1 {
			r1 := regexp.MustCompile(`:[a-zA-Z0-9_-]+`)
			s := r1.ReplaceAllString(item, `[a-zA-Z0-9_-]+`)
			r2 := regexp.MustCompile(s)
			checkLoginRegexps[i] = r2
		} else {
			checkLoginRegexps[i] = nil
		}
	}
}

var checkLoginUrls = []string{
	"/edit-article",
}

var checkLoginRegexps = make([]*regexp.Regexp, 10)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		url := c.Request.URL.String()
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
	for i, item := range checkLoginUrls {

		if strings.Index(item, "/:") > -1 {
			if checkLoginRegexps[i].MatchString(url) {
				return true
			}
		} else {
			if item == url {
				return true
			}
		}
	}
	return false
}
