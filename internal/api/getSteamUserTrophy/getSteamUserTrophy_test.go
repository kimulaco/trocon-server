package GetSteamUserTrophyAPI

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/h2non/gock"
	"github.com/kimulaco/trocon-server/pkg/httputil"
	"github.com/kimulaco/trocon-server/pkg/steamworks"
	"github.com/kimulaco/trocon-server/testdata"
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

func TestGetSteamUserTrophy_400(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerAchievementsPath,
		StatusCode: 400,
		Response: steamworks.GetPlayerAchievementsResponse{
			PlayerStats: steamworks.GetPlayerAchievementsResponseOwnedGame{
				GameName:     "",
				Achievements: []steamworks.Achievement{},
				Success: false,
			},
		},
	})

	rec, c := testdata.InitEcho("/api/steam/user/:steamid/trophy", "appid=1")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	if assert.NoError(t, GetSteamUserTrophy(c)) {
		assert.Equal(t, 200, rec.Code)

		resBody, err := httputil.ReadBody[SuccessResponse](rec.Result())
		assert.Equal(t, nil, err)

		expected := SuccessResponse{
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
		assert.Equal(t, expected, resBody)
	}
}

func TestGetSteamUserTrophy_403(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerAchievementsPath,
		StatusCode: 403,
		Response: testdata.SteamworksResponse403,
	})

	rec, c := testdata.InitEcho("/api/steam/user/:steamid/trophy", "appid=1")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	if assert.NoError(t, GetSteamUserTrophy(c)) {
		assert.Equal(t, 200, rec.Code)

		resBody, err := httputil.ReadBody[SuccessResponse](rec.Result())
		assert.Equal(t, nil, err)

		expected := SuccessResponse{
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
		assert.Equal(t, expected, resBody)
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
		Response: createPlayerAchievementsResponse(1),
	})

	rec, c := testdata.InitEcho("/api/steam/user/:steamid/trophy", "appid=1")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	if assert.NoError(t, GetSteamUserTrophy(c)) {
		assert.Equal(t, 200, rec.Code)
		resBody, err := httputil.ReadBody[SuccessResponse](rec.Result())
		assert.Equal(t, nil, err)
		assert.Equal(t, SuccessResponse{
			StatusCode: 200,
			Trophies: []Trophy{
				createDummyTrophy(1),
			},
		}, resBody)
	}
}

func TestGetSteamUserTrophy_TwoAppid(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	defer gock.Off()
	for i := 1; i <= 2; i++ {
		testdata.InitGock(testdata.GockConfig{
			Url:      "http://localhost:9999",
			Path:     steamworks.GetPlayerAchievementsPath,
			Querys:   map[string]string{"appid": strconv.Itoa(i)},
			Response: createPlayerAchievementsResponse(i),
		})
	}

	rec, c := testdata.InitEcho("/api/steam/user/:steamid/trophy", "appid=1,2")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	if assert.NoError(t, GetSteamUserTrophy(c)) {
		assert.Equal(t, 200, rec.Code)
		resBody, err := httputil.ReadBody[SuccessResponse](rec.Result())
		assert.Equal(t, nil, err)
		assert.Equal(t, SuccessResponse{
			StatusCode: 200,
			Trophies: []Trophy{
				createDummyTrophy(1),
				createDummyTrophy(2),
			},
		}, resBody)
	}
}

func BenchmarkGetSteamUserTrophy(b *testing.B) {
	os.Setenv("STEAM_API_KEY", "XXXXXXXX")
	os.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")
	defer os.Unsetenv("STEAM_API_KEY")
	defer os.Unsetenv("STEAM_API_BASE_URL")

	defer gock.Off()
	appidParams := ""
	for i := 1; i <= 21; i++ {
		appidParams += strconv.Itoa(i) + ","
		testdata.InitGock(testdata.GockConfig{
			Url:      "http://localhost:9999",
			Path:     steamworks.GetPlayerAchievementsPath,
			Querys:   map[string]string{"appid": strconv.Itoa(i)},
			Response: createPlayerAchievementsResponse(i),
		})
	}

	appidParams = appidParams[0 : len(appidParams) - 1]
	rec, c := testdata.InitEcho("/api/steam/user/:steamid/trophy", "appid=" + appidParams)
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := GetSteamUserTrophy(c); err != nil {
			b.Fatalf("Error during request: %s", err)
		}
		if rec.Code != http.StatusOK {
			b.Fatalf("Unexpected status code: %d", rec.Code)
		}
	}
}

func createDummyTrophy(appid int) Trophy {
	return Trophy{
		AppId:        appid,
		Success:      true,
		GameName:     "Trophy Game " + strconv.Itoa(appid),
		Achievements: []steamworks.Achievement{testdata.TestAchievement1, testdata.TestAchievement2},
	}
}

func createPlayerAchievementsResponse(appid int) steamworks.GetPlayerAchievementsResponse {
	return steamworks.GetPlayerAchievementsResponse{
		PlayerStats: steamworks.GetPlayerAchievementsResponseOwnedGame{
			GameName:     "Trophy Game " + strconv.Itoa(appid),
			Achievements: []steamworks.Achievement{testdata.TestAchievement1, testdata.TestAchievement2},
			Success: true,
		},
	}
}
