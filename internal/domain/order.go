package domain

import "time"

type Order struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Items     int
	CreatedAt time.Time
	Total     int
	Overshoot int
	Packs     []OrderPack `gorm:"foreignKey:OrderID"`
}

type OrderPack struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	OrderID uint
	Size    int
	Count   int
}
