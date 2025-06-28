package domain

type Pack struct {
	ID   uint `gorm:"primaryKey;autoIncrement"`
	Size int  `gorm:"uniqueIndex" json:"size"`
}
