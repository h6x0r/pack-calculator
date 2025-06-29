package dto

type PackAddRequest struct {
	Size int `json:"size" binding:"required,gt=0"`
}

type PackUpdateRequest struct {
	NewSize int `json:"new_size" binding:"required,gt=0"`
}

type CalculateRequest struct {
	Items int `form:"items" binding:"required,gte=0"`
}
