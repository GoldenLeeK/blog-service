package service

import (
	"github.com/GoldenLeeK/blog-service/interal/model"
	"github.com/GoldenLeeK/blog-service/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CreatedBy string `form:"created_by" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (s *Service) CountTag(param *CountTagRequest) (int, error) {
	return s.dao.CountTag(param.Name, param.State)
}

func (s *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return s.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (s *Service) CreateTag(param *CreateTagRequest) error {
	return s.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (s *Service) UpdateTag(param *UpdateTagRequest) error {
	return s.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (s *Service) DeleteTag(param *DeleteTagRequest) error {
	return s.dao.DeleteTag(param.ID)
}
