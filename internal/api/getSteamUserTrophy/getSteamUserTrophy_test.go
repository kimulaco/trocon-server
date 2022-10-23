package GetSteamUserTrophyAPI

import (
	"errors"
	"testing"

	"github.com/h2non/gock"
	"github.com/kimulaco/trophy-comp-server/pkg/httputil"
	"github.com/kimulaco/trophy-comp-server/pkg/steamworks"
	"github.com/kimulaco/trophy-comp-server/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetSteamUserTrophy_EnvError(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	rec, c := testdata.InitEcho("/api/steam/user/:steamid/trophy", "")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	if assert.NoError(t, GetSteamUserTrophy(c)) {
		assert.Equal(t, 500, rec.Code)
		resBody, err := httputil.ReadBody[httputil.Error](rec.Result())
		assert.Equal(t, nil, err)
		assert.Equal(t, httputil.NewError(
			500,
			"STEAM_USER_TROPHY_ENVIROMENT_ERROR",
			errors.New("internal server error"),
		), resBody)
	}
}

func TestGetSteamUserTrophy_SteamidNotFound(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	rec, c := testdata.InitEcho("/api/steam/user/:steamid/trophy", "")
	c.SetParamNames("steamid")
	c.SetParamValues("")

	if assert.NoError(t, GetSteamUserTrophy(c)) {
		assert.Equal(t, 400, rec.Code)
		resBody, err := httputil.ReadBody[httputil.Error](rec.Result())
		assert.Equal(t, nil, err)
		assert.Equal(t, httputil.NewError(
			400,
			"STEAM_USER_TROPHY_STEAMID_NOT_FOUND",
			errors.New("steamid not found"),
		), resBody)
	}
}

func TestGetSteamUserTrophy_AppidNotFound(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	rec, c := testdata.InitEcho("/api/steam/user/:steamid/trophy", "")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	if assert.NoError(t, GetSteamUserTrophy(c)) {
		assert.Equal(t, 400, rec.Code)
		resBody, err := httputil.ReadBody[httputil.Error](rec.Result())
		assert.Equal(t, nil, err)
		assert.Equal(t, httputil.NewError(
			400,
			"STEAM_USER_TROPHY_APPID_NOT_FOUND",
			errors.New("appid not found"),
		), resBody)
	}
}

func TestGetSteamUserTrophy_AppNotFound(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerAchievementsPath,
		Response: GetSteamUserTrophy_AppNotFound,
	})

	rec, c := testdata.InitEcho("/api/steam/user/:steamid/trophy", "appid=1")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	if assert.NoError(t, GetSteamUserTrophy(c)) {
		assert.Equal(t, 200, rec.Code)
		resBody, err := httputil.ReadBody[SuccessResponse](rec.Result())
		assert.Equal(t, nil, err)
		assert.Equal(t, SuccessResponse_AppNotFound, resBody)
	}
}

func TestGetSteamUserTrophy_OneAppid(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerAchievementsPath,
		Querys:   map[string]string{"appid": "1"},
		Response: GetSteamUserTrophy_Appid1,
	})

	rec, c := testdata.InitEcho("/api/steam/user/:steamid/trophy", "appid=1")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	if assert.NoError(t, GetSteamUserTrophy(c)) {
		assert.Equal(t, 200, rec.Code)
		resBody, err := httputil.ReadBody[SuccessResponse](rec.Result())
		assert.Equal(t, nil, err)
		assert.Equal(t, SuccessResponse_OneAppid, resBody)
	}
}

func TestGetSteamUserTrophy_TwoAppid(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerAchievementsPath,
		Querys:   map[string]string{"appid": "1"},
		Response: GetSteamUserTrophy_Appid1,
	})
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerAchievementsPath,
		Querys:   map[string]string{"appid": "2"},
		Response: GetSteamUserTrophy_Appid2,
	})

	rec, c := testdata.InitEcho("/api/steam/user/:steamid/trophy", "appid=1,2")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	if assert.NoError(t, GetSteamUserTrophy(c)) {
		assert.Equal(t, 200, rec.Code)
		resBody, err := httputil.ReadBody[SuccessResponse](rec.Result())
		assert.Equal(t, nil, err)
		assert.Equal(t, SuccessResponse_TwoAppid, resBody)
	}
}

var SuccessResponse_AppNotFound = SuccessResponse{
	StatusCode: 200,
	Trophies: []Trophy{
		{
			AppId:        1,
			Success:      false,
			GameName:     "",
			Achievements: []steamworks.Achievement{},
		},
	},
}

var SuccessResponse_OneAppid = SuccessResponse{
	StatusCode: 200,
	Trophies: []Trophy{
		{
			AppId:        1,
			Success:      true,
			GameName:     "Trophy Game 1",
			Achievements: []steamworks.Achievement{testdata.TestAchievement1, testdata.TestAchievement2},
		},
	},
}

var SuccessResponse_TwoAppid = SuccessResponse{
	StatusCode: 200,
	Trophies: []Trophy{
		{
			AppId:        1,
			Success:      true,
			GameName:     "Trophy Game 1",
			Achievements: []steamworks.Achievement{testdata.TestAchievement1, testdata.TestAchievement2},
		},
		{
			AppId:        2,
			Success:      true,
			GameName:     "Trophy Game 2",
			Achievements: []steamworks.Achievement{testdata.TestAchievement1, testdata.TestAchievement2},
		},
	},
}

var GetSteamUserTrophy_AppNotFound = steamworks.GetPlayerAchievementsResponse{
	PlayerStats: steamworks.GetPlayerAchievementsResponseOwnedGame{
		GameName:     "",
		Achievements: make([]steamworks.Achievement, 0),
	},
}

var GetSteamUserTrophy_Appid1 = steamworks.GetPlayerAchievementsResponse{
	PlayerStats: steamworks.GetPlayerAchievementsResponseOwnedGame{
		GameName:     "Trophy Game 1",
		Achievements: []steamworks.Achievement{testdata.TestAchievement1, testdata.TestAchievement2},
	},
}

var GetSteamUserTrophy_Appid2 = steamworks.GetPlayerAchievementsResponse{
	PlayerStats: steamworks.GetPlayerAchievementsResponseOwnedGame{
		GameName:     "Trophy Game 2",
		Achievements: []steamworks.Achievement{testdata.TestAchievement1, testdata.TestAchievement2},
	},
}
