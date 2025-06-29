package calc

import (
	"errors"
	"github.com/h6x0r/pack-calculator/internal/application/calc/dto"
	"math"
)

func Calculate(order int, sizes []int) (dto.CalculateResponse, error) {
	if order < 0 {
		return dto.CalculateResponse{}, errors.New("order must be >= 0")
	}
	if len(sizes) == 0 {
		return dto.CalculateResponse{}, errors.New("no pack sizes provided")
	}

	minPacks := sizes[0]
	for _, s := range sizes {
		if s < minPacks {
			minPacks = s
		}
	}

	limit := order + minPacks
	const inf = math.MaxInt32

	dp := make([]int, limit+1)
	prev := make([]int, limit+1)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0

	for _, p := range sizes {
		for sum := 0; sum+p <= limit; sum++ {
			if dp[sum] != inf && dp[sum]+1 < dp[sum+p] {
				dp[sum+p] = dp[sum] + 1
				prev[sum+p] = p
			}
		}
	}

	target := -1
	for t := order; t <= limit; t++ {
		if dp[t] != inf {
			target = t
			break
		}
	}
	if target == -1 {
		return dto.CalculateResponse{}, errors.New("cannot fulfill order")
	}

	packs := map[int]int{}
	for t := target; t > 0; {
		p := prev[t]
		packs[p]++
		t -= p
	}

	return dto.CalculateResponse{
		Packs:     packs,
		Total:     target,
		Overshoot: target - order,
	}, nil
}
