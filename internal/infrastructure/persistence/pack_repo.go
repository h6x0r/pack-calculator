package persistence

import (
	"github.com/h6x0r/pack-calculator/internal/domain"
	"github.com/h6x0r/pack-calculator/internal/infrastructure/persistence/dto"
	"gorm.io/gorm"
)

type PackRepo struct{ db *gorm.DB }

func NewPackRepo(db *gorm.DB) *PackRepo { return &PackRepo{db} }

func (r *PackRepo) List() ([]domain.Pack, error) {
	var entities []dto.PackEntity
	if err := r.db.Order("size").Find(&entities).Error; err != nil {
		return nil, err
	}

	packs := make([]domain.Pack, len(entities))
	for i, entity := range entities {
		packs[i] = MapEntityToPack(entity)
	}

	return packs, nil
}

func (r *PackRepo) Create(size int) (domain.Pack, error) {
	entity := dto.PackEntity{Size: size}
	if err := r.db.Create(&entity).Error; err != nil {
		return domain.Pack{}, err
	}

	return MapEntityToPack(entity), nil
}

func (r *PackRepo) Delete(size int) error {
	return r.db.Where("size = ?", size).Delete(&dto.PackEntity{}).Error
}

func (r *PackRepo) Update(oldSize, newSize int) error {
	return r.db.Model(&dto.PackEntity{}).
		Where("size = ?", oldSize).
		Update("size", newSize).Error
}
