package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gin-contrib/multitemplate"
)

func getFilesList(path, stuffix string) (files []string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		// 将模板文件放到列表
		if strings.HasSuffix(info.Name(), stuffix) {
			files = append(files, path)
		}
		return nil
	})
	return
}

// go test -v -run TestTemplate ./test/template_test.go
func TestTemplate(t *testing.T) {
	templateDir := "../templates"
	stuffix := ".html"
	r := multitemplate.NewRenderer()
	// 查找模板路径下所有文件
	rd, _ := ioutil.ReadDir(templateDir)
	fmt.Printf("rd: %v\n", rd)
	for _, fi := range rd {
		if fi.IsDir() {
			// 文件夹
			for _, f := range getFilesList(path.Join(templateDir, fi.Name()), stuffix) {
				fmt.Printf("f: %v\n", f)
				fmt.Printf("f[len(templateDir)+1:len(f)-len(stuffix)]: %v\n", f[len(templateDir)+1:len(f)-len(stuffix)])
				r.AddFromFiles(f[len(templateDir)+1:len(f)-len(stuffix)], f)
			}

		} else {
			// 模板文件，直接添加
			t.Logf("file.Name(): %v", fi.Name())
			if strings.HasSuffix(fi.Name(), stuffix) {
				t.Log(path.Join(templateDir, fi.Name()))
				r.AddFromFiles(path.Join(templateDir, fi.Name()))
			}
		}
	}
}
