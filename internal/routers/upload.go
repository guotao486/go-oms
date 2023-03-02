package routers

import (
	"oms/global"
	"oms/internal/service"
	"oms/pkg/app"
	"oms/pkg/convert"
	"oms/pkg/errcode"
	"oms/pkg/upload"

	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

// Upload/file godoc
// @Tags Upload
// @Summary 文件上传
// @Produce json
// @Param type formData int true "上传类型" default(1)
// @Param file formData file true "file"
// @Success 200 {object} app.Success "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /upload/file [post]
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	// 获取上传的文件对象
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	// 判断上传类型参数合法性
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	success := app.NewSuccess()
	success.Data = gin.H{
		"file_access_url": fileInfo.AccessUrl,
	}
	response.ToResponse(success)
	return
}
