package config

import (
	"os"
)

func GetListenPort(defaultPort string) string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	return ":" + port
}
