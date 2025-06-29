package dto

type PackAddRequest struct {
	Size int
}

type PackUpdateRequest struct {
	OldSize int
	NewSize int
}

type PackDeleteRequest struct {
	Size int
}
