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
	r.HTMLRender = LoadTemplateFiles("./views", ".html")
	InitControllers(r)

	r.Run(":8080")
}

func LoadTemplateFiles(templateDir, suffix string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	rd, _ := ioutil.ReadDir(templateDir)
	for _, fi := range rd {
		if fi.IsDir() {
			for _, f := range getFileList(path.Join(templateDir, fi.Name()), suffix) {
				println(f[len(templateDir)-1:len(f)-len(suffix)])
				println(f)
				r.AddFromFiles(f[len(templateDir)-1:len(f)-len(suffix)], f)
				//r.AddFromFiles(f[len(templateDir)-1:], f)
			}
		} else {
			r.AddFromFiles(fi.Name(), path.Join(templateDir, fi.Name()))
		}
	}

	fmt.Println(r)

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
