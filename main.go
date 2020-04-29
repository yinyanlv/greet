package main

import (
	"fmt"
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

	r.Static("/static", "./static")
	r.HTMLRender = loadTemplates("./views")
	//r.HTMLRender = LoadTemplateFiles("./views", ".html")
	InitControllers(r)

	r.Run(":1124")
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

func LoadTemplateFiles(templateDir, suffix string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	rd, _ := ioutil.ReadDir(templateDir)
	for _, fi := range rd {
		if fi.IsDir() {
			for _, f := range getFileList(path.Join(templateDir, fi.Name()), suffix) {
				r.AddFromFiles(f[len(templateDir) + 1:], f)
			}
		} else {
			r.AddFromFiles(fi.Name(), path.Join(templateDir, fi.Name()))
		}
	}

	fmt.Println(r)
	return r
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/common/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}