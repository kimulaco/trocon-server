package steamworks

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestGetPlayerAchievements(t *testing.T) {
	const baseUrl = "http://localhost:9999"
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", baseUrl)
	s := NewSteamworks()

	defer gock.Off()
	gock.New(baseUrl).
		Get(GetPlayerAchievementsPath).
		Reply(200).
		JSON(GetPlayerAchievementsResponse200)

	Game, err := s.GetPlayerAchievements("1", "1")
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, Game, GetPlayerAchievementsResponse200.PlayerStats)
}

var GetPlayerAchievementsResponse200 = GetPlayerAchievementsResponse{
	PlayerStats: GetPlayerAchievementsResponseOwnedGame{
		GameName:     "Trophy Game 1",
		Achievements: []Achievement{TestAchievement1, TestAchievement2},
	},
}

var TestAchievement1 = Achievement{
	ApiName:     "api-1",
	Name:        "Trophy 1",
	Description: "",
	Achieved:    1,
	UnlockTime:  1640962800,
}

var TestAchievement2 = Achievement{
	ApiName:     "api-2",
	Name:        "Trophy 2",
	Description: "",
	Achieved:    0,
	UnlockTime:  1640962800,
}
