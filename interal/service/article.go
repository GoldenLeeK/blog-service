package service

import (
	"github.com/GoldenLeeK/blog-service/interal/model"
	"github.com/GoldenLeeK/blog-service/pkg/app"
)

type CountArticleRequest struct {
	Title string `form:"title" binding:"max=128"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=128"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string   `form:"title" binding:"max=128,required"`
	Desc          string   `form:"desc" binding:"max=512,required"`
	Content       string   `form:"content" binding:"required"`
	CoverImageUrl string   `form:"cover_image_url" binding:"required,url"`
	CreatedBy     string   `form:"created_by" binding:"max=255,required"`
	State         uint8    `form:"state,default=1" binding:"oneof=0 1"`
	Tags          []uint32 `form:"tags" binding:"required"`
}

func (s *Service) CountArticle(param *CountArticleRequest) (int, error) {
	return s.dao.CountArticle(param.Title, param.State)
}

func (s *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	return s.dao.GetArticleList(param.Title, param.State, pager.Page, pager.PageSize)
}

func (s *Service) CreateArticle(param *CreateArticleRequest) error {
	return s.dao.CreateArticle(param.Title, param.Desc, param.Content, param.CoverImageUrl, param.CreatedBy, param.State, param.Tags)
}
