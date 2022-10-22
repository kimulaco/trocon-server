package steamworks

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestGetOwnedGames(t *testing.T) {
	const baseUrl = "http://localhost:9999"
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", baseUrl)
	s := NewSteamworks()

	defer gock.Off()
	gock.New(baseUrl).
		Get(GetOwnedGamesPath).
		Reply(200).
		JSON(GetOwnedGamesResponse200)

	ownedGames, err := s.GetOwnedGames(TestUser.SteamID)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, ownedGames, GetOwnedGamesResponse200.Response.Games)
}

var GetOwnedGamesResponse200 = GetOwnedGamesResponse{
	Response: struct {
		GameCount int         "json:\"game_count\""
		Games     []OwnedGame "json:\"games\""
	}{
		GameCount: 2,
		Games:     []OwnedGame{TestGame1, TestGame2},
	},
}
