package config

import (
	"log"
	"os"
	"strings"
)

type Config struct {
	HTTPPort string
	DBPath   string
}

func Load() Config {
	get := func(key, def string) string {
		if v := strings.TrimSpace(os.Getenv(key)); v != "" {
			return v
		}
		return def
	}

	cfg := Config{
		HTTPPort: get("PACK_CALC_PORT", ":8081"),
		DBPath:   get("PACK_CAL_DB", "pack_calc.db"),
	}
	log.Printf("[config] %+v", cfg)
	return cfg
}
