package v1

import (
	"github.com/GoldenLeeK/blog-service/global"
	"github.com/GoldenLeeK/blog-service/interal/service"
	"github.com/GoldenLeeK/blog-service/pkg/app"
	"github.com/GoldenLeeK/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func NewTag() *Tag {
	return &Tag{}
}

// @Summary 获取多个标签
// @Produce  json
// @Tags 标签组
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t *Tag) List(c *gin.Context) {
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &service.TagListRequest{})

	if !valid {
		global.Logger.Errorf("app.BindAndValid errs : %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	response.ToResponse(gin.H{})

	return
}

// @Summary 新增标签
// @Produce  json
// @Tags 标签组
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t *Tag) Create(c *gin.Context) {}

// @Summary 更新标签
// @Produce  json
// @Tags 标签组
// @Param id path int true "标签 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t *Tag) Update(c *gin.Context) {}

// @Summary 删除标签
// @Produce  json
// @Tags 标签组
// @Param id path int true "标签 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t *Tag) Delete(c *gin.Context) {}
