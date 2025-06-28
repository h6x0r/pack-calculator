package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"pc/internal/application/calc"
	"pc/internal/application/pack"
)

type Handlers struct {
	calcSvc *calc.Service
	packSvc *pack.Service
}

func (h *Handlers) PacksList(c *gin.Context) {
	ps, _ := h.packSvc.List()
	c.JSON(http.StatusOK, ps)
}

func (h *Handlers) PacksAdd(c *gin.Context) {
	var b struct {
		Size int `json:"size"`
	}
	if err := c.ShouldBindJSON(&b); err != nil || b.Size <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "size must be >0"})
		return
	}
	if _, err := h.packSvc.Add(b.Size); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handlers) PacksDelete(c *gin.Context) {
	size, err := strconv.Atoi(c.Param("size"))
	if err != nil || size <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid size"})
		return
	}
	if err := h.packSvc.Remove(size); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handlers) PacksUpdate(c *gin.Context) {
	oldSize, err := strconv.Atoi(c.Param("size"))
	if err != nil || oldSize <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid size"})
		return
	}
	var b struct {
		NewSize int `json:"new_size"`
	}
	if err := c.ShouldBindJSON(&b); err != nil || b.NewSize <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "new_size must be >0"})
		return
	}
	if err := h.packSvc.Change(oldSize, b.NewSize); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handlers) CalculateGet(c *gin.Context) {
	items, err := strconv.Atoi(c.Query("items"))
	if err != nil || items < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "items param required"})
		return
	}
	h.respondCalc(c, items)
}

func (h *Handlers) respondCalc(c *gin.Context, items int) {
	res, err := h.calcSvc.Calculate(items)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
