package persistence

import (
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"pc/config"
	"pc/internal/domain"
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
		if err := db.AutoMigrate(&domain.Pack{}, &domain.Order{}, &domain.OrderPack{}); err != nil {
			log.Fatal(err)
		}
		var cnt int64
		db.Model(&domain.Pack{}).Count(&cnt)
		if cnt == 0 {
			for _, s := range []int{250, 500, 1000, 2000, 5000} {
				db.Create(&domain.Pack{Size: s})
			}
		}
	})
	return db
}
