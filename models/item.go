package models

import "time"

type Item struct {
	ItemId      uint			`gorm:"primaryKey" json:"itemId"`
	ItemCode    string		`gorm:"type:varchar(100)" json:"itemCode" binding:"required"`
	Description string		`gorm:"type:text" json:"description" binding:"required"`
	Quantity    int				`json:"quantity" binding:"required"`
	OrderId     uint			`json:"orderId"`
	CreatedAt   time.Time	`json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}