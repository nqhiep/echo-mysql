package model

import (
	"time"
)

type User struct {
	Id          string     `json:"id" gorm:"column:id"`
	Username    string     `json:"username" gorm:"column:username"`
	Email       string     `json:"email" gorm:"column:email"`
	Phone       string     `json:"phone" gorm:"column:phone"`
	DateOfBirth *time.Time `json:"dateOfBirth" gorm:"column:date_of_birth"`
}
