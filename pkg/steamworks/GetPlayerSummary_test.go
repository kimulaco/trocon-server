package steamworks

import (
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

var TestUser = Player{
	SteamID:                  "0",
	CommunityVisibilityState: 3,
	ProfileState:             1,
	PersonaName:              "trophy-comp-user",
	LastLogoff:               1640962800,
	ProfileUrl:               "http://localhost:9999/id/0/",
	Avatar:                   "http://localhost:9999/avatar.jpg",
	AvatarMedium:             "http://localhost:9999/avatar_medium.jpg",
	AvatarFull:               "http://localhost:9999/avatar_full.jpg",
}

var GetPlayerSummaryResponse200 = GetPlayerSummaryResponse{
	Response: struct {
		Players []Player "json:\"players\""
	}{
		Players: []Player{TestUser},
	},
}
