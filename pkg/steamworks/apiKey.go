package steamworks

import (
	"os"
)

type ApiKey struct {
	Key string
}

func (t ApiKey) HasKey() bool {
	return t.Key == ""
}

func NewApiKey() ApiKey {
	STEAM_API_KEY := os.Getenv("STEAM_API_KEY")

	return ApiKey{
		Key: STEAM_API_KEY,
	}
}
