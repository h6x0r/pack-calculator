package api

import (
	calc "github.com/h6x0r/pack-calculator/internal/application/calc/dto"
	pack "github.com/h6x0r/pack-calculator/internal/application/pack/dto"
	api "github.com/h6x0r/pack-calculator/internal/infrastructure/api/dto"
)

func MapCalculateRequest(req api.CalculateRequest) calc.CalculateRequest {
	return calc.CalculateRequest{
		Items: req.Items,
	}
}

func MapCalculateResponse(resp calc.CalculateResponse) api.CalculateResponse {
	return api.CalculateResponse{
		Packs:     resp.Packs,
		Total:     resp.Total,
		Overshoot: resp.Overshoot,
	}
}

func MapPackAddRequest(req api.PackAddRequest) pack.PackAddRequest {
	return pack.PackAddRequest{
		Size: req.Size,
	}
}

func MapPackResponse(resp pack.PackResponse) api.PackResponse {
	return api.PackResponse{
		ID:   resp.ID,
		Size: resp.Size,
	}
}

func MapPackUpdateRequest(oldSize int, req api.PackUpdateRequest) pack.PackUpdateRequest {
	return pack.PackUpdateRequest{
		OldSize: oldSize,
		NewSize: req.NewSize,
	}
}

func MapPackDeleteRequest(size int) pack.PackDeleteRequest {
	return pack.PackDeleteRequest{
		Size: size,
	}
}
