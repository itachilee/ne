package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// MyModel
	Id              uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Name            string `json:"name" gorm:"column:name"`
	Email           string `json:"email" gorm:"column:email"`
	Status          int    `json:"status" gorm:"column:status"`
	EmailVerifiedAt int64  `json:"email_verified_at" gorm:"column:email_verified_at"`
	Password        string `json:"password" gorm:"column:password"`
	RememberToken   string `json:"remember_token" gorm:"column:remember_token"`
}

func (t User) TableName() string {
	return "z_users"
}

func (t User) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (v MyModel) BeforeCreate(tx *gorm.DB) error {

	tx.Statement.SetColumn("created_at", time.Now())
	tx.Statement.SetColumn("updated_at", time.Now())
	return nil
}

func (t User) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("status = ?", t.Status)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t User) List(db *gorm.DB, pageOffset, pageSize int) ([]*User, error) {
	var tags []*User
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	if t.Email != "" {
		db = db.Where("email =?", t.Email)
	}
	db = db.Where("status = ?", t.Status)
	if err = db.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t User) Update(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := db.Model(&User{}).Where("id = ?", t.Id).Update("status", t.Status).Error; err != nil {
			return err
		}
		return nil
	})

	// 条件更新
	return err
}
