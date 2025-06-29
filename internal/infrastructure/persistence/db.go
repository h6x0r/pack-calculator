package persistence

import (
	"log"
	"sync"

	"github.com/h6x0r/pack-calculator/config"
	"github.com/h6x0r/pack-calculator/internal/infrastructure/persistence/dto"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func DB(cfg config.Config) *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(cfg.DBPath+"?_busy_timeout=10000"), &gorm.Config{})
		if err != nil {
			log.Fatal("open db:", err)
		}
		if err := db.AutoMigrate(&dto.PackEntity{}, &dto.OrderEntity{}, &dto.OrderPackEntity{}); err != nil {
			log.Fatal(err)
		}
		var cnt int64
		db.Model(&dto.PackEntity{}).Count(&cnt)
		if cnt == 0 {
			for _, s := range []int{250, 500, 1000, 2000, 5000} {
				db.Create(&dto.PackEntity{Size: s})
			}
		}
	})
	return db
}
