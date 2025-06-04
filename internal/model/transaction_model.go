package model

import "time"

type Transaction struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      User      `gorm:"foreignKey:UserID; references:ID"`
	Type      string    `gorm:"type:enum('in','out'); not null; default:'in'"`
	CreatedAt time.Time `gorm:"type:timestamp; not null; default:current_timestamp"`
	Item      []TransactionItem
	DeletedAt time.Time `gorm:"index"`
}
