/*
 * @Author: GG
 * @Date: 2023-01-31 17:04:44
 * @LastEditTime: 2023-02-01 16:14:39
 * @LastEditors: GG
 * @Description: 文件操作
 * @FilePath: \oms\pkg\upload\file.go
 *
 */
package upload

import (
	"io"
	"mime/multipart"
	"oms/global"
	"oms/pkg/util"
	"os"
	"path"
	"strings"
)

// 类型别名
type FileType int

const TypeImage FileType = iota + 1

func GetFileName(name string) string {
	ext := GetFileExt(name)
	// strings.TrimSuffix 返回没有提供的尾随后缀字符串的 s。如果 s 不以 suffix 结尾，则 s 原样返回
	filename := strings.TrimSuffix(name, ext)
	filename = util.EncodeMD5(filename)
	return filename + ext
}

/**
 * @description: 返回文件扩展名
 * @param {string} name
 * @return {*}
 */
func GetFileExt(name string) string {
	// path.Ext 返回文件扩展名
	return path.Ext(name)
}

/**
 * @description: 获取文件保存地址
 * @return {*}
 */
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// ====== 文件检查部分 =======

/**
 * @description: 检查目录是否存在
 * @param {string} dst
 * @return {*}
 */
func CheckSavePath(dst string) bool {
	// os.Stat 方法获取文件的描述信息 FileInfo
	_, err := os.Stat(dst)
	// os.IsNotExist 检查文件错误信息是否为文件或目录不存在
	return os.IsNotExist(err)
}

/**
 * @description: 检查文件权限
 * @param {string} dst
 * @return {*}
 */
func CheckPermission(dst string) bool {
	// os.Stat 方法获取文件的描述信息 FileInfo
	_, err := os.Stat(dst)
	// os.IsPermission 检查文件错误信息是否为权限不足
	return os.IsPermission(err)
}

/**
 * @description: 检查文件类型和后缀是否匹配
 * @param {FileType} t
 * @param {string} name
 * @return {*}
 */
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	// 后缀统一转大写匹配
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			// 后缀统一转大写匹配
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

/**
 * @description: 检查文件大小是否超出最大大小限制
 * @param {FileType} t
 * @param {multipart.File} f
 * @return {*}
 */
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := io.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size <= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// ====== 创建和写入 ======

// CreateSavePath
//
// 创建在上传文件时所使用的保存目录，在方法内部调用的 os.MkdirAll 方法，该方法将会以传入的 os.FileMode 权限位去递归创建所需的所有目录结构，若涉及的目录均已存在，则不会进行任何操作，直接返回 nil。
//
/**
 * @description: 创建文件目录
 * @param {string} dst
 * @param {os.FileMode} perm
 * @return {*}
 */
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

// SaveFile
// 保存所上传的文件，该方法主要是通过调用 os.Create 方法创建目标地址的文件，再通过 file.Open 方法打开源地址的文件，结合 io.Copy 方法实现两者之间的文件内容拷贝
//
/**
 * @description: 保存上传文件
 * @param {*multipart.FileHeader} file
 * @param {string} dst
 * @return {*}
 */
func SaveFile(file *multipart.FileHeader, dst string) error {
	// 从给定的FileHeader读信息，返回二维数组
	// 打开源文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 创建目标地址文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer out.Close()

	// 将接收的文件内容copy到新创建的文件
	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}
	return nil
}
