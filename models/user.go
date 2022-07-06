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

func (t User) Count(db *gorm.DB) *int64 {
	umgr := ZUsersMgr(db)
	var count *int64
	umgr.WithName(t.Name)
	umgr.WithStatus(int8(t.Status))
	_ = umgr.Count(count)

	return count
}

func (t User) List(db *gorm.DB, pageOffset, pageSize int) (interface{}, error) {

	umgr := ZUsersMgr(db)
	// 构造分页条件
	page := NewPage((int64(pageOffset)), int64(pageOffset))
	// 不带条件的查询
	result, err := umgr.SelectPage(page, umgr.WithStatus(1))
	if err != nil {
		return nil, err
	}
	// result.GetPages()
	return result.GetRecords(), nil
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
