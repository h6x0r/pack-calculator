package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/h6x0r/pack-calculator/internal/application/calc"
	"github.com/h6x0r/pack-calculator/internal/application/pack"
	apidto "github.com/h6x0r/pack-calculator/internal/infrastructure/api/dto"
)

type Handlers struct {
	calcSvc *calc.Service
	packSvc *pack.Service
}

func (h *Handlers) PacksList(c *gin.Context) {
	response, err := h.packSvc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, apidto.ErrorResponse{Error: err.Error()})
		return
	}

	apiResponse := make([]apidto.PackResponse, len(response.Packs))
	for i, p := range response.Packs {
		apiResponse[i] = MapPackResponse(p)
	}

	c.JSON(http.StatusOK, apiResponse)
}

func (h *Handlers) PacksAdd(c *gin.Context) {
	var req apidto.PackAddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, apidto.ErrorResponse{Error: "size must be >0"})
		return
	}

	serviceReq := MapPackAddRequest(req)
	response, err := h.packSvc.Add(serviceReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, apidto.ErrorResponse{Error: err.Error()})
		return
	}

	apiResponse := MapPackResponse(response)
	c.JSON(http.StatusCreated, apiResponse)
}

func (h *Handlers) PacksDelete(c *gin.Context) {
	size, err := strconv.Atoi(c.Param("size"))
	if err != nil || size <= 0 {
		c.JSON(http.StatusBadRequest, apidto.ErrorResponse{Error: "invalid size"})
		return
	}

	serviceReq := MapPackDeleteRequest(size)
	if err := h.packSvc.Remove(serviceReq); err != nil {
		c.JSON(http.StatusBadRequest, apidto.ErrorResponse{Error: err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handlers) PacksUpdate(c *gin.Context) {
	oldSize, err := strconv.Atoi(c.Param("size"))
	if err != nil || oldSize <= 0 {
		c.JSON(http.StatusBadRequest, apidto.ErrorResponse{Error: "invalid size"})
		return
	}

	var req apidto.PackUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, apidto.ErrorResponse{Error: "new_size must be >0"})
		return
	}

	serviceReq := MapPackUpdateRequest(oldSize, req)
	if err := h.packSvc.Change(serviceReq); err != nil {
		c.JSON(http.StatusBadRequest, apidto.ErrorResponse{Error: err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handlers) CalculateGet(c *gin.Context) {
	items, err := strconv.Atoi(c.Query("items"))
	if err != nil || items < 0 {
		c.JSON(http.StatusBadRequest, apidto.ErrorResponse{Error: "items param required"})
		return
	}
	h.respondCalc(c, items)
}

func (h *Handlers) respondCalc(c *gin.Context, items int) {
	apiReq := apidto.CalculateRequest{Items: items}
	serviceReq := MapCalculateRequest(apiReq)
	res, err := h.calcSvc.Calculate(serviceReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, apidto.ErrorResponse{Error: err.Error()})
		return
	}

	apiResponse := MapCalculateResponse(res)
	c.JSON(http.StatusOK, apiResponse)
}
