package model

import "gorm.io/gorm"

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100); not null"`
	Description string `gorm:"type:text; not null"`
	ImagePath   string `gorm:"type:varchar(255); not null"`
	CategoryID  uint
	Category    Category       `gorm:"foreignKey:CategoryID; references:ID"`
	Stock       uint           `gorm:"not null; default:0"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
