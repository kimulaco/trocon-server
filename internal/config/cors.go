package config

import (
	"log"
	"os"
)

const MODE_PROD = "production"
const MODE_DEV = "development"

func GetCORSAllowOrigins() []string {
	mode := os.Getenv("BUILD_MODE")

	if mode == MODE_PROD {
		origin := os.Getenv("CORS_ALLOW_ORIGIN")

		if origin != "" {
			return []string{origin}
		}

		log.Print("CORS_ALLOW_ORIGIN not found")
		return []string{}
	}

	if mode == MODE_DEV {
		return []string{"http://localhost:*"}
	}

	return []string{}
}
