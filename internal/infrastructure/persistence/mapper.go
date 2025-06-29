package persistence

import (
	"github.com/h6x0r/pack-calculator/internal/domain"
	"github.com/h6x0r/pack-calculator/internal/infrastructure/persistence/dto"
)

func MapEntityToPack(entity dto.PackEntity) domain.Pack {
	return domain.Pack{
		ID:   entity.ID,
		Size: entity.Size,
	}
}

func MapOrderToEntity(order domain.Order) dto.OrderEntity {
	entity := dto.OrderEntity{
		Items:     order.Items,
		Total:     order.Total,
		Overshoot: order.Overshoot,
	}

	orderPacks := make([]dto.OrderPackEntity, len(order.Packs))
	for i, pack := range order.Packs {
		orderPacks[i] = dto.OrderPackEntity{
			Size:  pack.Size,
			Count: pack.Count,
		}
	}
	entity.Packs = orderPacks

	return entity
}
