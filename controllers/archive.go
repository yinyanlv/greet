package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prot/models"
	"strconv"
)

func RenderArchive(c *gin.Context) {
	pageIndex, _ := strconv.ParseUint(c.Query("page"), 10, 64)
	pageSize, _ := strconv.ParseUint(c.Query("size"), 10, 64)
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}
	a := models.Article{}
	var articles []models.Article
	var err error

	articles, err = a.GetArticlesByTag("", pageIndex, pageSize)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"errorCode": 599,
			"message":   err,
		})
		return
	}

	c.HTML(http.StatusOK, "list", gin.H{
		"pageCode":  "archive",
		"pageTitle": "归档",
		"articles":  articles,
	})
}

func RenderArchiveByYearMonth(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")
	pageIndex, _ := strconv.ParseUint(c.Query("page"), 10, 64)
	pageSize, _ := strconv.ParseUint(c.Query("size"), 10, 64)
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}
	a := models.Article{}
	var articles []models.Article
	var err error

	articles, err = a.GetArticlesByYearMonth(year, month, pageIndex, pageSize)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"errorCode": 599,
			"message":   err,
		})
		return
	}

	c.HTML(http.StatusOK, "list", gin.H{
		"pageCode":  "archive",
		"pageTitle": "归档 " + year + "-" + month,
		"articles":  articles,
	})
}
