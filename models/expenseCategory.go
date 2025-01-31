package models

import "time"

type ExpenseCategory struct {
	ID        uint      `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name"`
	UserID    uint      `gorm:"colum:nuserid;unique"`
	CreatedAt time.Time `gorm:"createdat;autoCreateTime"`
	UpdatedAt time.Time `gorm:"updatedat;autoCreateTime"`
}
