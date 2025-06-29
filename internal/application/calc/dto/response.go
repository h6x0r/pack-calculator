package dto

type CalculateResponse struct {
	Packs     map[int]int
	Total     int
	Overshoot int
}
