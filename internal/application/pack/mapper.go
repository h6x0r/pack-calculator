package pack

import (
	"github.com/h6x0r/pack-calculator/internal/application/pack/dto"
	"github.com/h6x0r/pack-calculator/internal/domain"
)

func MapDomainToPackResponse(pack domain.Pack) dto.PackResponse {
	return dto.PackResponse{
		ID:   pack.ID,
		Size: pack.Size,
	}
}

func MapDomainListToPackListResponse(packs []domain.Pack) dto.PackListResponse {
	packResponses := make([]dto.PackResponse, len(packs))
	for i, pack := range packs {
		packResponses[i] = MapDomainToPackResponse(pack)
	}
	return dto.PackListResponse{Packs: packResponses}
}
