package utils

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
)

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
