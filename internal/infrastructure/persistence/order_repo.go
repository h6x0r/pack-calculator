package persistence

import (
	"gorm.io/gorm"
	"pc/internal/domain"
)

type OrderRepo struct{ db *gorm.DB }

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{db}
}

func (r *OrderRepo) Save(o *domain.Order) error {
	return r.db.Create(o).Error
}
