package steamworks

import (
	"os"
)

type Steamworks struct {
	APIKey     string
	APIBaseURL string
}

func (s Steamworks) InvalidEnv() bool {
	return s.APIKey == "" || s.APIBaseURL == ""
}

func NewSteamworks() Steamworks {
	key := os.Getenv("STEAM_API_KEY")
	url := os.Getenv("STEAM_API_BASE_URL")

	return Steamworks{
		APIKey:     key,
		APIBaseURL: url,
	}
}
