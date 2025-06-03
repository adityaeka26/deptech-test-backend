package model

import "gorm.io/gorm"

type Category struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"type:varchar(100); not null"`
	Description string         `gorm:"type:text; not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
