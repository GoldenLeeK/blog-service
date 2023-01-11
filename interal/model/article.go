package model

import (
	"github.com/GoldenLeeK/blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title" gorm:"index;type:varchar(128)"`
	Desc          string `json:"desc" gorm:"type:varchar(512)"`
	Content       string `json:"content" gorm:"type:text;"`
	CoverImageUrl string `json:"cover_image_url" gorm:"type:text"`
	Tags          []*Tag `gorm:"many2many:article_tags"`
	State         uint8  `json:"state" gorm:"type:tinyint;"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a *Article) Count(db *gorm.DB) (int, error) {
	var count int
	if a.Title != "" {
		db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)
	if err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (a *Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	var err error

	db = db.Preload("Tags")
	if pageOffset > 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}

	db = db.Where("state = ?", a.State)

	if err = db.Where("is_del = ?", 0).Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, err

}
