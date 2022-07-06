package dto

type UserDto struct {
	Name            string `json:"name" gorm:"column:name"`
	Email           string `json:"email" gorm:"column:email"`
	Status          int    `json:"status" gorm:"column:status"`
	EmailVerifiedAt int64  `json:"email_verified_at" gorm:"column:email_verified_at"`
}
