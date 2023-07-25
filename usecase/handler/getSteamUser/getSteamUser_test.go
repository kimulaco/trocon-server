package GetSteamUserAPI

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/h2non/gock"
	Game "github.com/kimulaco/trocon-server/domain/game"
	"github.com/kimulaco/trocon-server/interface/steamworks"
	"github.com/kimulaco/trocon-server/pkg/httputil"
	"github.com/kimulaco/trocon-server/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetSteamUser_EnvError(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	rec, c := testdata.InitEcho("/api/steam/user/:steamid", "")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	if assert.NoError(t, GetSteamUser(c)) {
		resBody, err := httputil.ReadBody[httputil.Error](rec.Result())
		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, nil, err)
		assert.Equal(t, httputil.NewError(
			500,
			"STEAM_USER_ENVIROMENT_ERROR",
			errors.New("internal server error"),
		), resBody)
	}
}

func TestGetSteamUser_SteamidNotFound(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	rec, c := testdata.InitEcho("/api/steam/user/:steamid", "")
	c.SetParamNames("steamid")
	c.SetParamValues("")

	if assert.NoError(t, GetSteamUser(c)) {
		resBody, err := httputil.ReadBody[httputil.Error](rec.Result())
		assert.Equal(t, 400, rec.Code)
		assert.Equal(t, nil, err)
		assert.Equal(t, httputil.NewError(
			400,
			"STEAM_USER_STEAMID_NOT_FOUND",
			errors.New("steamid not found"),
		), resBody)
	}
}

func TestGetSteamUser_PlayerNotFound(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerSummaryPath,
		Response: GetPlayerSummaryResponse_Empty,
	})
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetOwnedGamesPath,
		Response: GetOwnedGamesResponse_Empty,
	})

	rec, c := testdata.InitEcho("/api/steam/user/:steamid", "")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	if assert.NoError(t, GetSteamUser(c)) {
		resBody, err := httputil.ReadBody[httputil.Error](rec.Result())
		assert.Equal(t, 404, rec.Code)
		assert.Equal(t, nil, err)
		assert.Equal(t, httputil.NewError(
			404,
			"STEAM_USER_NOT_FOUND",
			errors.New("user not found"),
		), resBody)
	}
}

func TestGetSteamUser_FailGetPlayer(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:        "http://localhost:9999",
		Path:       steamworks.GetPlayerSummaryPath,
		StatusCode: 500,
		Response:   GetPlayerSummaryResponse_Empty,
	})
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetOwnedGamesPath,
		Response: GetOwnedGamesResponse_200,
	})

	rec, c := testdata.InitEcho("/api/steam/user/:steamid", "")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	if assert.NoError(t, GetSteamUser(c)) {
		resBody, err := httputil.ReadBody[httputil.Error](rec.Result())
		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, nil, err)
		assert.Equal(t, httputil.NewError(
			500,
			"STEAM_USER_INTERNAL_ERROR",
			errors.New("internal server error"),
		), resBody)
	}
}

func TestGetSteamUser_FailGetGames(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerSummaryPath,
		Response: GetPlayerSummaryResponse_200,
	})
	testdata.InitGock(testdata.GockConfig{
		Url:        "http://localhost:9999",
		Path:       steamworks.GetOwnedGamesPath,
		StatusCode: 500,
		Response:   GetOwnedGamesResponse_Empty,
	})

	rec, c := testdata.InitEcho("/api/steam/user/:steamid", "")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	if assert.NoError(t, GetSteamUser(c)) {
		resBody, err := httputil.ReadBody[httputil.Error](rec.Result())
		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, nil, err)
		assert.Equal(t, httputil.NewError(
			500,
			"STEAM_OWNED_GAME_INTERNAL_ERROR",
			errors.New("internal server error"),
		), resBody)
	}
}

func TestGetSteamUser_200(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerSummaryPath,
		Response: GetPlayerSummaryResponse_200,
	})
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetOwnedGamesPath,
		Response: GetOwnedGamesResponse_200,
	})

	rec, c := testdata.InitEcho("/api/steam/user/:steamid", "")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	if assert.NoError(t, GetSteamUser(c)) {
		resBody, err := httputil.ReadBody[GetSteamUserResponse](rec.Result())
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, nil, err)
		assert.Equal(t, SuccessResponse_200, resBody)
	}
}

func BenchmarkGetSteamUser_200(b *testing.B) {
	os.Setenv("STEAM_API_KEY", "XXXXXXXX")
	os.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")
	defer os.Unsetenv("STEAM_API_KEY")
	defer os.Unsetenv("STEAM_API_BASE_URL")

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerSummaryPath,
		Response: GetPlayerSummaryResponse_200,
	})
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetOwnedGamesPath,
		Response: GetOwnedGamesResponse_200,
	})

	rec, c := testdata.InitEcho("/api/steam/user/:steamid", "")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	for i := 0; i < b.N; i++ {
		if err := GetSteamUser(c); err != nil {
			b.Fatalf("Error during request: %s", err)
		}
		if rec.Code != http.StatusOK {
			b.Fatalf("Unexpected status code: %d", rec.Code)
		}
	}
}

var GetPlayerSummaryResponse_200 = steamworks.GetPlayerSummaryResponse{
	Response: struct {
		Players []steamworks.Player "json:\"players\""
	}{
		Players: []steamworks.Player{steamworks.TestUser},
	},
}

var GetPlayerSummaryResponse_Empty = steamworks.GetPlayerSummaryResponse{
	Response: struct {
		Players []steamworks.Player "json:\"players\""
	}{
		Players: make([]steamworks.Player, 0),
	},
}

var GetOwnedGamesResponse_200 = steamworks.GetOwnedGamesResponse{
	Response: struct {
		GameCount int                    "json:\"game_count\""
		Games     []steamworks.OwnedGame "json:\"games\""
	}{
		GameCount: 2,
		Games:     []steamworks.OwnedGame{steamworks.TestGame1, steamworks.TestGame2},
	},
}

var GetOwnedGamesResponse_Empty = steamworks.GetOwnedGamesResponse{
	Response: struct {
		GameCount int                    "json:\"game_count\""
		Games     []steamworks.OwnedGame "json:\"games\""
	}{
		GameCount: 0,
		Games:     make([]steamworks.OwnedGame, 0),
	},
}

var SuccessResponse_200 = GetSteamUserResponse{
	StatusCode: 200,
	User:       steamworks.TestUser,
	Games:      []Game.Game{steamworks.TestGame1.ToGame(), steamworks.TestGame2.ToGame()},
}
