package app

import (
	"html/template"
	"io/ioutil"
	"oms/global"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
)

func toLinux(basePath string) string {
	return strings.ReplaceAll(basePath, "\\", "/")
}
func getFilesList(path, stuffix string) (files []string) {
	// walk win环境的地址为\，需要将其转换一下
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		// 将模板文件放到列表
		if strings.HasSuffix(info.Name(), stuffix) {
			files = append(files, toLinux(path))
		}
		return nil
	})
	return
}

// 加载指定路径全部模板文件
func LoadTemplateFiles() multitemplate.Renderer {
	templateDir, templateLayoutDir, stuffix := global.AppSetting.TemplatePath, global.AppSetting.TemplateLayoutPath, global.AppSetting.TemplateStuffix
	ajaxTemplate := "ajaxForm"
	r := multitemplate.NewRenderer()
	// 查找模板路径下所有文件
	rd, _ := ioutil.ReadDir(templateDir)

	var layoutFiles []string   // 根模板文件
	var moduleFiles []string   // 模块模板
	var templateFiles []string // 根目录模板
	var ajaxFormFiles []string
	for _, fi := range rd {
		if fi.IsDir() {
			files := getFilesList(path.Join(templateDir, fi.Name()), stuffix)
			// 根模板文件
			if fi.Name() == templateLayoutDir {
				layoutFiles = append(layoutFiles, files...)
			} else if fi.Name() == ajaxTemplate {
				ajaxFormFiles = append(ajaxFormFiles, files...)
			} else {
				// 模块模板
				moduleFiles = append(moduleFiles, files...)
			}

		} else {
			// 模板文件，直接添加
			if strings.HasSuffix(fi.Name(), stuffix) {
				templateFiles = append(templateFiles, fi.Name())
			}
		}
	}

	// 根目录模板处理
	for _, f := range templateFiles {
		r.AddFromFiles(f[:len(f)-len(stuffix)], path.Join(templateDir, f))
	}
	for _, f := range ajaxFormFiles {
		r.AddFromFiles(f[len(templateDir)+1:len(f)-len(stuffix)], f)
	}
	// 模块模板处理
	for _, f := range moduleFiles {
		r.AddFromFilesFuncs(f[len(templateDir)+1:len(f)-len(stuffix)], LoadFuncs(), append(layoutFiles, f)...)
	}

	return r
}

// 加载自定义函数
func LoadFuncs() template.FuncMap {
	return template.FuncMap{
		"getImagesUrl": GetImagesUrl,
		"inString":     InString,
	}
}

// 返回完整的图片地址
func GetImagesUrl(images string) string {
	images = global.AppSetting.UploadServerUrl + "/" + images

	return images
}

// 检查拼接的字符串是否有出现的关键词
func InString(s, k string) bool {
	ss := strings.Split(s, ",")
	for _, v := range ss {
		if k == v {
			return true
		}
	}
	return false
}
