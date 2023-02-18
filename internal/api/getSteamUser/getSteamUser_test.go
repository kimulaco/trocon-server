package GetSteamUserAPI

import (
	"errors"
	"testing"

	"github.com/h2non/gock"
	"github.com/kimulaco/trocon-server/pkg/httputil"
	"github.com/kimulaco/trocon-server/pkg/steamworks"
	"github.com/kimulaco/trocon-server/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetUser_EnvError(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	rec, c := testdata.InitEcho("/api/steam/user/:steamid", "")
	c.SetParamNames("steamid")
	c.SetParamValues("1")

	if assert.NoError(t, GetUser(c)) {
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

func TestGetUser_SteamidNotFound(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	rec, c := testdata.InitEcho("/api/steam/user/:steamid", "")
	c.SetParamNames("steamid")
	c.SetParamValues("")

	if assert.NoError(t, GetUser(c)) {
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

func TestGetUser_PlayerNotFound(t *testing.T) {
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

	if assert.NoError(t, GetUser(c)) {
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

func TestGetUser_200(t *testing.T) {
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

	if assert.NoError(t, GetUser(c)) {
		resBody, err := httputil.ReadBody[GetUserResponse](rec.Result())
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, nil, err)
		assert.Equal(t, SuccessResponse_200, resBody)
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

var SuccessResponse_200 = GetUserResponse{
	StatusCode: 200,
	User:       steamworks.TestUser,
	Games:      []steamworks.Game{steamworks.TestGame1.ToGame(), steamworks.TestGame2.ToGame()},
}
