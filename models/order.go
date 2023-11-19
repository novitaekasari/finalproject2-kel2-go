package models

import "time"

type Order struct {
	OrderId      uint				`gorm:"primaryKey" json:"orderId"`
	CustomerName string			`gorm:"type:varchar(100)" json:"customerName" binding:"required"`
	OrderedAt    time.Time	`json:"orderedAt" binding:"required"`
	CreatedAt    time.Time	`json:"createdAt"`
	UpdatedAt    time.Time	`json:"updatedAt"`
	Items        []Item			`json:"items" binding:"required"`
}