package steamworks

import (
	"errors"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestGetPlayerAchievements_200(t *testing.T) {
	const baseUrl = "http://localhost:9999"
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", baseUrl)
	s := NewSteamworks()

	defer gock.Off()
	r := gock.New(baseUrl)
	r.Get(GetPlayerAchievementsPath).Reply(200).JSON(GetPlayerAchievementsResponse200)

	Game, err := s.GetPlayerAchievements("1", "1")
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, Game, GetPlayerAchievementsResponse200.PlayerStats)
}

func TestGetPlayerAchievements_400(t *testing.T) {
	const baseUrl = "http://localhost:9999"
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", baseUrl)
	s := NewSteamworks()

	defer gock.Off()
	r := gock.New(baseUrl)
	r.Get(GetPlayerAchievementsPath).Reply(400).JSON(GetPlayerAchievementsResponse400)

	Game, err := s.GetPlayerAchievements("1", "1")
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, Game, GetPlayerAchievementsResponse400.PlayerStats)
}

func TestGetPlayerAchievements_Error(t *testing.T) {
	const baseUrl = "http://localhost:9999"
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", baseUrl)
	s := NewSteamworks()

	defer gock.Off()
	r := gock.New(baseUrl)

	r.Get(GetPlayerAchievementsPath).Reply(200).JSON(GetPlayerAchievementsResponse200)
	res1, err1 := s.GetPlayerAchievements("", "1")
	assert.Equal(t, res1, EmptyOwnedGame)
	assert.Equal(t, err1, errors.New("steamid is required"))

	r.Get(GetPlayerAchievementsPath).Reply(200).JSON(GetPlayerAchievementsResponse200)
	res2, err2 := s.GetPlayerAchievements("1", "")
	assert.Equal(t, res2, EmptyOwnedGame)
	assert.Equal(t, err2, errors.New("appid is required"))

	r.Get(GetPlayerAchievementsPath).Reply(403).JSON(Response403)
	res3, err3 := s.GetPlayerAchievements("1", "1")
	assert.Equal(t, res3, EmptyOwnedGame)
	assert.Equal(t, err3, errors.New("403 Forbidden"))
}

var EmptyOwnedGame = GetPlayerAchievementsResponseOwnedGame{
	GameName: "",
	Achievements: make([]Achievement, 0),
	Error: "",
	Success: false,
}

var GetPlayerAchievementsResponse200 = GetPlayerAchievementsResponse{
	PlayerStats: GetPlayerAchievementsResponseOwnedGame{
		GameName:     "Trophy Game 1",
		Achievements: []Achievement{TestAchievement1, TestAchievement2},
		Success: true,
	},
}

var GetPlayerAchievementsResponse400 = GetPlayerAchievementsResponse{
	PlayerStats: GetPlayerAchievementsResponseOwnedGame{
		GameName: "",
		Achievements: []Achievement{},
		Error:     "Requested app has no stats",
		Success: false,
	},
}
