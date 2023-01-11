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

func (d *Dao) CreateArticle(title, desc, content, CoverImageUrl, createdBy string, state uint8, tagIds []uint32) error {
	tag := &model.Tag{}
	tags, _ := tag.InTags(d.engine, tagIds)

	article := model.Article{
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: CoverImageUrl,
		Tags:          tags,
		State:         state,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}

	return article.Create(d.engine)
}
