package dto

import "time"

type PackEntity struct {
	ID   uint `gorm:"primaryKey;autoIncrement"`
	Size int  `gorm:"uniqueIndex"`
}

type OrderEntity struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Items     int
	CreatedAt time.Time
	Total     int
	Overshoot int
	Packs     []OrderPackEntity `gorm:"foreignKey:OrderID"`
}

type OrderPackEntity struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	OrderID uint
	Size    int
	Count   int
}
