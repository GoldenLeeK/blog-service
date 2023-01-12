package v1

import (
	"github.com/GoldenLeeK/blog-service/global"
	"github.com/GoldenLeeK/blog-service/interal/service"
	"github.com/GoldenLeeK/blog-service/pkg/app"
	"github.com/GoldenLeeK/blog-service/pkg/convert"
	"github.com/GoldenLeeK/blog-service/pkg/errcode"
	"github.com/GoldenLeeK/blog-service/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct {
}

func NewUpload() *Upload {
	return &Upload{}
}

func (u *Upload) UploadFile(c *gin.Context) {
	respone := app.Response{Ctx: c}
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		respone.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		respone.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		respone.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	respone.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
