package calc

import (
	"errors"
	"github.com/h6x0r/pack-calculator/internal/application/calc/dto"
	"github.com/h6x0r/pack-calculator/internal/domain"
)

type Service struct {
	packs  domain.PackRepository
	orders domain.OrderRepository
}

func New(p domain.PackRepository, o domain.OrderRepository) *Service {
	return &Service{packs: p, orders: o}
}

func (s *Service) Calculate(req dto.CalculateRequest) (dto.CalculateResponse, error) {
	if req.Items < 0 {
		return dto.CalculateResponse{}, errors.New("items must be ≥0")
	}

	ps, err := s.packs.List()
	if err != nil {
		return dto.CalculateResponse{}, err
	}
	if len(ps) == 0 {
		return dto.CalculateResponse{}, errors.New("no pack sizes configured")
	}

	sizes := make([]int, len(ps))
	for i, p := range ps {
		sizes[i] = p.Size
	}

	res, err := Calculate(req.Items, sizes)
	if err != nil {
		return dto.CalculateResponse{}, err
	}

	order := domain.Order{Items: req.Items, Total: res.Total, Overshoot: res.Overshoot}
	for size, cnt := range res.Packs {
		order.Packs = append(order.Packs, domain.OrderPack{Size: size, Count: cnt})
	}
	if err := s.orders.Save(&order); err != nil {
		return dto.CalculateResponse{}, err
	}

	return res, nil
}
