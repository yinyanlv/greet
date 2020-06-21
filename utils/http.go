package utils

import (
	"github.com/gin-gonic/gin"
)

func IsAjax(c *gin.Context) bool {
	key := "x-requested-with"
	val := "XMLHttpRequest"
	if c.Request.Header.Get(key) == val {
		return true
	} else {
		return false
	}
}
