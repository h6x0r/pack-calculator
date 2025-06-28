package persistence

import (
	"gorm.io/gorm"
	"pc/internal/domain"
)

type PackRepo struct{ db *gorm.DB }

func NewPackRepo(db *gorm.DB) *PackRepo { return &PackRepo{db} }

func (r *PackRepo) List() ([]domain.Pack, error) {
	var ps []domain.Pack
	return ps, r.db.Order("size").Find(&ps).Error
}

func (r *PackRepo) Create(size int) (domain.Pack, error) {
	p := domain.Pack{Size: size}
	return p, r.db.Create(&p).Error
}

func (r *PackRepo) Delete(size int) error {
	return r.db.Where("size = ?", size).Delete(&domain.Pack{}).Error
}

func (r *PackRepo) Update(oldSize, newSize int) error {
	return r.db.Model(&domain.Pack{}).
		Where("size = ?", oldSize).
		Update("size", newSize).Error
}
