package calc

import (
	"errors"
	"pc/internal/domain"
)

type Service struct {
	packs  domain.PackRepository
	orders domain.OrderRepository
}

func New(p domain.PackRepository, o domain.OrderRepository) *Service {
	return &Service{packs: p, orders: o}
}

type Result struct {
	Packs     map[int]int `json:"packs"`
	Total     int         `json:"total"`
	Overshoot int         `json:"overshoot"`
}

func (s *Service) Calculate(items int) (Result, error) {
	if items < 0 {
		return Result{}, errors.New("items must be ≥0")
	}

	ps, err := s.packs.List()
	if err != nil {
		return Result{}, err
	}
	if len(ps) == 0 {
		return Result{}, errors.New("no pack sizes configured")
	}

	sizes := make([]int, len(ps))
	for i, p := range ps {
		sizes[i] = p.Size
	}

	res, err := Calculate(items, sizes)
	if err != nil {
		return Result{}, err
	}

	order := domain.Order{Items: items, Total: res.Total, Overshoot: res.Overshoot}
	for size, cnt := range res.Packs {
		order.Packs = append(order.Packs, domain.OrderPack{Size: size, Count: cnt})
	}
	if err := s.orders.Save(&order); err != nil {
		return Result{}, err
	}
	return res, nil
}
