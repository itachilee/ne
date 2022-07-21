package repository

import (
	"log"

	"github.com/itachilee/gin/models"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

// UserRepository : represent the user's repository contract
type UserRepository interface {
	AddUser(models.ZUsers) (models.ZUsers, error)
	GetUser(int) (models.ZUsers, error)
	GetByEmail(string) (models.ZUsers, error)
	GetAllUser() ([]models.ZUsers, error)
	UpdateUser(models.ZUsers) (models.ZUsers, error)
	DeleteUser(models.ZUsers) (models.ZUsers, error)
	Migrate() error
}

// NewUserRepository -> returns new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) Migrate() error {
	log.Print("[UserRepository]...Migrate")
	return u.DB.AutoMigrate(&models.ZUsers{})
}

func (u userRepository) GetUser(id int) (user models.ZUsers, err error) {
	return user, u.DB.First(&user, id).Error
}

func (u userRepository) GetByEmail(email string) (user models.ZUsers, err error) {
	return user, u.DB.First(&user, "email=?", email).Error
}

func (u userRepository) GetAllUser() (users []models.ZUsers, err error) {
	return users, u.DB.Find(&users).Error
}

func (u userRepository) AddUser(user models.ZUsers) (models.ZUsers, error) {
	return user, u.DB.Create(&user).Error
}

func (u userRepository) UpdateUser(user models.ZUsers) (models.ZUsers, error) {
	if err := u.DB.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, u.DB.Model(&user).Updates(&user).Error
}

func (u userRepository) DeleteUser(user models.ZUsers) (models.ZUsers, error) {
	if err := u.DB.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, u.DB.Delete(&user).Error
}
