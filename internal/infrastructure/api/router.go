package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pc/internal/application/calc"
	"pc/internal/application/pack"
	"pc/internal/infrastructure/persistence"
)

func Router(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.StaticFile("/", "./ui/index.html")

	packRepo := persistence.NewPackRepo(db)
	orderRepo := persistence.NewOrderRepo(db)

	h := &Handlers{
		calcSvc: calc.New(packRepo, orderRepo),
		packSvc: pack.New(packRepo),
	}

	api := r.Group("/api/v1")
	{
		api.GET("/packs", h.PacksList)
		api.POST("/packs", h.PacksAdd)
		api.PUT("/packs/:size", h.PacksUpdate)
		api.DELETE("/packs/:size", h.PacksDelete)
		api.GET("/calculate", h.CalculateGet)
	}

	return r
}
