package model

import "gorm.io/gorm"

type TransactionItem struct {
	ID            uint `gorm:"primaryKey"`
	TransactionID uint
	Transaction   Transaction `gorm:"foreignKey:TransactionID; references:ID"`
	ProductID     uint
	Product       Product `gorm:"foreignKey:ProductID; references:ID"`
	Quantity      uint
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
