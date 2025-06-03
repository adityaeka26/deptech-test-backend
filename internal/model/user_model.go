package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primaryKey"`
	FirstName   string         `gorm:"type:varchar(50); not null"`
	LastName    string         `gorm:"type:varchar(50); not null"`
	Email       string         `gorm:"type:varchar(255); uniqueIndex; not null"`
	Password    string         `gorm:"type:varchar(255); not null"`
	DateOfBirth time.Time      `gorm:"type:date; not null"`
	Gender      string         `gorm:"type:enum('M','F'); not null; default:'M'"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
