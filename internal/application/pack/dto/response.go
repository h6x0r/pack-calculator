package dto

type PackResponse struct {
	ID   uint
	Size int
}

type PackListResponse struct {
	Packs []PackResponse
}
