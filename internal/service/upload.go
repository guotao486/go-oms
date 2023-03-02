package service

import (
	"errors"
	"mime/multipart"
	"oms/global"
	"oms/pkg/upload"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

// UploadFile
//
/**
 * @description: 文件上传验证与处理
 * @param {upload.FileType} fileType
 * @param {multipart.File} file
 * @param {*multipart.FileHeader} fileHeader
 * @return {*}
 */
func (s *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	filename := upload.GetFileName(fileHeader.Filename)

	// 检查文件后缀
	if !upload.CheckContainExt(fileType, filename) {
		return nil, errors.New("file suffix is not supported.")
	}

	// 检查文件大小
	if !upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}

	// 先检查是否存在目录，不存在则创建
	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		// 创建目录
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	// 检查是否有保存文件权限
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}

	dst := uploadSavePath + "/" + filename
	// 保存文件
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + filename
	return &FileInfo{Name: filename, AccessUrl: accessUrl}, nil
}
