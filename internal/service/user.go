package service

import (
	"github.com/itachilee/gin/models"
	"github.com/itachilee/gin/pkg/app"
)

type CreateUserRequest struct {
	Name  string `form:"name" binding:"min=6,max=100"`
	Email string `form:"email" binding:"min=6" `

	Password string `form:"password" binding:"min=6" `
}

type UserListRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status int    `form:"status,default=1" binding:"oneof=0 1"`
}

type UpdateUserRequest struct {
	Id     uint64 `form:"id" binding:"min=1"`
	Status int    `form:"status,default=1" binding:"oneof=0 1"`
}

func (svc *Service) CreateUser(param *CreateUserRequest) error {
	return svc.dao.CreateUser(param.Name, param.Email, param.Password)
}

func (svc *Service) CountUser(param *UserListRequest) (int64, error) {

	return svc.dao.CountUser(param.Name, param.Status)
}

func (svc *Service) GetUserList(param *UserListRequest, pager *app.Pager) ([]*models.User, error) {
	return svc.dao.GetUserList(param.Name, param.Status, pager.Page, pager.PageSize)
}

func (svc *Service) UpdateUser(param *UpdateUserRequest) error {
	return svc.dao.UpdateUser(param.Id, param.Status)
}
