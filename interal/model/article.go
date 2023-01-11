package model

import "github.com/GoldenLeeK/blog-service/pkg/app"

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
