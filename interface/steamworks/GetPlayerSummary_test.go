package steamworks

import (
	"errors"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestGetPlayerSummary(t *testing.T) {
	const baseUrl = "http://localhost:9999"
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", baseUrl)
	s := NewSteamworks()

	defer gock.Off()
	gock.New(baseUrl).
		Get(GetPlayerSummaryPath).
		Reply(200).
		JSON(GetPlayerSummaryResponse200)

	player, err := s.GetPlayerSummary(TestUser.SteamID)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, player, TestUser)
}

func TestGetPlayerSummaryError(t *testing.T) {
	const baseUrl = "http://localhost:9999"
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", baseUrl)
	s := NewSteamworks()

	defer gock.Off()
	r := gock.New(baseUrl)

	r.Get(GetPlayerSummaryPath).Reply(200).JSON(GetPlayerSummaryResponse200)
	player1, err1 := s.GetPlayerSummary("")
	assert.Equal(t, player1, Player{})
	assert.Equal(t, err1, errors.New("steamid is required"))

	r.Get(GetPlayerSummaryPath).Reply(403).JSON(Response403)
	player2, err2 := s.GetPlayerSummary("1")
	assert.Equal(t, player2, Player{SteamID: "1"})
	assert.Equal(t, err2, errors.New("403 Forbidden"))
}

var GetPlayerSummaryResponse200 = GetPlayerSummaryResponse{
	Response: struct {
		Players []Player "json:\"players\""
	}{
		Players: []Player{TestUser},
	},
}
