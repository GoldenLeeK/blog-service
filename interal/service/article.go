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

func (s *Service) CountArticle(param *CountArticleRequest) (int, error) {
	return s.dao.CountArticle(param.Title, param.State)
}

func (s *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	return s.dao.GetArticleList(param.Title, param.State, pager.Page, pager.PageSize)
}
