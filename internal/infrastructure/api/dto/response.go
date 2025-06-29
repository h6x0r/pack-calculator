package dto

type PackResponse struct {
	ID   uint `json:"id"`
	Size int  `json:"size"`
}

type CalculateResponse struct {
	Packs     map[int]int `json:"packs"`
	Total     int         `json:"total"`
	Overshoot int         `json:"overshoot"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
