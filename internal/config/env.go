package config

import (
	"os"
)

const MODE_PROD = "production"
const MODE_DEV = "development"

func GetCORSAllowOrigins() []string {
	mode := os.Getenv("BUILD_MODE")

	if mode == MODE_PROD {
		return []string{}
	}

	if mode == MODE_DEV {
		return []string{"http://localhost:*"}
	}

	return []string{}
}
