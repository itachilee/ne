package dao

import (
	"github.com/itachilee/gin/models"
	"github.com/itachilee/gin/pkg/app"
)

func (d *Dao) CreateUser(name, email, password string) error {
	user := models.User{
		Name:            name,
		Email:           email,
		Password:        password,
		Status:          1,
		EmailVerifiedAt: 0_0,
	}

	return user.Create(d.engine)
}

func (d *Dao) CountUser(name string, status int) (int64, error) {
	user := models.User{Name: name, Status: status}
	return user.Count(d.engine)
}

func (d *Dao) GetUserList(name string, status int, page, pageSize int) ([]*models.User, error) {
	tag := models.User{Name: name, Status: status}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) UpdateUser(id uint64, status int) error {
	user := models.User{Id: id, Status: status}
	return user.Update(d.engine)
}
