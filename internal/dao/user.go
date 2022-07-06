package dao

import (
	"github.com/itachilee/gin/internal/dto"
	"github.com/itachilee/gin/models"
	"github.com/itachilee/gin/pkg/app"
	"github.com/jinzhu/copier"
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

func (d *Dao) CountUser(name string, status int) *int64 {
	user := models.User{Name: name, Status: status}
	return user.Count(d.engine)
}

func (d *Dao) GetUserList(name string, status int, page, pageSize int) (interface{}, error) {
	tag := models.User{Name: name, Status: status}
	pageOffset := app.GetPageOffset(page, pageSize)
	v, err := tag.List(d.engine, pageOffset, pageSize)
	if err != nil {
		return nil, err
	}
	var users []models.ZUsers
	if _, ok := v.([]models.ZUsers); ok {
		// fmt.Println("int64")
		users = v.([]models.ZUsers)
	}
	var dtos []dto.UserDto
	copier.Copy(&dtos, &users)
	return dtos, nil
}

func (d *Dao) UpdateUser(id uint64, status int) error {
	user := models.User{Id: id, Status: status}
	return user.Update(d.engine)
}
