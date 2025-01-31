package models

import "time"

type User struct {
	ID        uint      `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	IsActive  bool      `gorm:"column:password"`
	CreatedAt time.Time `gorm:"createdat;autoCreateTime"`
	UpdatedAt time.Time `gorm:"updatedat;autoCreateTime"`
	DeletedAt time.Time `gorm:"delatedat;autoCreateTime"`
}
