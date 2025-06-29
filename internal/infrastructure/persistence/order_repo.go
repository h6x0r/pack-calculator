package persistence

import (
	"github.com/h6x0r/pack-calculator/internal/domain"
	"gorm.io/gorm"
)

type OrderRepo struct{ db *gorm.DB }

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{db}
}

func (r *OrderRepo) Save(o *domain.Order) error {
	entity := MapOrderToEntity(*o)
	return r.db.Create(&entity).Error
}
