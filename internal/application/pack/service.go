package pack

import (
	"github.com/h6x0r/pack-calculator/internal/application/pack/dto"
	"github.com/h6x0r/pack-calculator/internal/domain"
)

type Service interface {
	List() (dto.PackListResponse, error)
	Add(req dto.PackAddRequest) (dto.PackResponse, error)
	Remove(req dto.PackDeleteRequest) error
	Change(req dto.PackUpdateRequest) error
}

type ServiceImpl struct {
	repo domain.PackRepository
}

func New(r domain.PackRepository) *ServiceImpl {
	return &ServiceImpl{r}
}

func (s *ServiceImpl) List() (dto.PackListResponse, error) {
	packs, err := s.repo.List()
	if err != nil {
		return dto.PackListResponse{}, err
	}

	return MapDomainListToPackListResponse(packs), nil
}

func (s *ServiceImpl) Add(req dto.PackAddRequest) (dto.PackResponse, error) {
	pack, err := s.repo.Create(req.Size)
	if err != nil {
		return dto.PackResponse{}, err
	}

	return MapDomainToPackResponse(pack), nil
}

func (s *ServiceImpl) Remove(req dto.PackDeleteRequest) error {
	return s.repo.Delete(req.Size)
}

func (s *ServiceImpl) Change(req dto.PackUpdateRequest) error {
	return s.repo.Update(req.OldSize, req.NewSize)
}
