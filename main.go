package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	. "prot/src/db"
)

func main() {

	connStr := "root:111111@tcp(127.0.0.1:3306)/prot?charset=utf8&parseTime=True&loc=Local"
	ConnDb(connStr)
	Migrate()

	r := gin.Default()
	commonTpls := getFileList("./views/common", ".html")

	r.Static("/static", "./static")
	r.HTMLRender = LoadTemplateFiles("./views", ".html", commonTpls)
	InitControllers(r)

	r.Run(":8080")
}

func LoadTemplateFiles(templateDir, suffix string, commonTpls []string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	rd, _ := ioutil.ReadDir(templateDir)
	for _, fi := range rd {
		if !fi.IsDir() {
			arr := make([]string, 0)
			filename := fi.Name()
			p := path.Join(templateDir, filename)
			tpls := append(append(arr, p), commonTpls...)
			r.AddFromFiles(filename[0:len(filename)-len(suffix)], tpls...)
		}
	}
	return r
}

func getFileList(path string, suffix string) (files []string) {

	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {

		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, suffix) {
			files = append(files, path)
		}

		return nil
	})

	return files
}
