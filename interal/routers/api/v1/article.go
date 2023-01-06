package v1

import (
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() *Article {
	return &Article{}
}

// @Summary 获取单个文章
// @Produce  json
// @Tags 文章组
// @Param id path int true "文章id"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {}

// @Summary 获取多个文章
// @Produce  json
// @Tags 文章组
// @Param title query string false "文章标题" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {}

// @Summary 新增文章
// @Produce  json
// @Tags 文章组
// @Param title body string true "文章名称" minlength(1) maxlength(100)
// @Param desc body string true "文章简称" minlength(1) maxlength(255)
// @Param content body string true "文章内容" minlength(1)
// @Param cover_image_url body string true "文章封面链接" minlength(1)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {}

// @Summary 更新文章
// @Produce  json
// @Tags 文章组
// @Param id path int true "文章id"
// @Param title body string true "文章名称" minlength(1) maxlength(100)
// @Param desc body string true "文章简称" minlength(1) maxlength(255)
// @Param content body string true "文章内容" minlength(1)
// @Param cover_image_url body string true "文章封面链接" minlength(1)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {}

// @Summary 删除文章
// @Produce  json
// @Tags 文章组
// @Param id path int true "文章 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {}
