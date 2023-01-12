package model

import "github.com/jinzhu/gorm"

type Auth struct {
	*Model
	AppKey    string `json:"app_key" gorm:"type:varchar(20)"`
	AppSecret string `json:"app_secret" gorm:"type:varchar(50)"`
}

func (a *Auth) Get(db *gorm.DB) (*Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.AppKey, a.AppSecret, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &auth, err
	}

	return &auth, nil
}
