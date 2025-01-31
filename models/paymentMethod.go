package models

import "time"

type PaymentMethod struct {
	ID        uint      `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name"`
	UserID    uint      `gorm:"columnuserid;unique"`
	CreatedAt time.Time `gorm:"createdat;autoCreateTime"`
	UpdatedAt time.Time `gorm:"updatedat;autoCreateTime"`
}
