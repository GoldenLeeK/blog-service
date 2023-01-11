package dao

import (
	"github.com/GoldenLeeK/blog-service/interal/model"
	"github.com/GoldenLeeK/blog-service/pkg/app"
)

func (d *Dao) GetArticleList(title string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := &model.Article{Title: title, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CountArticle(title string, state uint8) (int, error) {
	article := model.Article{Title: title, State: state}
	return article.Count(d.engine)
}
