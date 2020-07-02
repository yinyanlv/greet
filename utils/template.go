package utils

import (
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"prot/models"
	"strings"
	"time"

	"github.com/gin-contrib/multitemplate"
)

var funcMap = template.FuncMap{
	"formatDatetime": func(dt time.Time) string {
		return dt.Format("2006-01-02 15:04:05")
	},
	"year": func(dt time.Time) int {
		return dt.Year()
	},
	"month": func(dt time.Time) int {
		return int(dt.Month())
	},
	"isActiveTag": func(id string, tags []models.Tag) bool {
		for _, item := range tags {
			if item.ID == id {
				return true
			}
		}
		return false
	},
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
			r.AddFromFilesFuncs(filename[0:len(filename)-len(suffix)], funcMap, tpls...)
		}
	}
	return r
}

func GetFileList(path string, suffix string) (files []string) {

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
