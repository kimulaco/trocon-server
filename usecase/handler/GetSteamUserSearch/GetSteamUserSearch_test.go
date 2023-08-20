package GetSteamUserSearchAPI

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/kimulaco/trocon-server/interface/steamworks"
	"github.com/kimulaco/trocon-server/pkg/httputil"
	"github.com/kimulaco/trocon-server/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetUser_MatchSteamID(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerSummaryPath,
		Response: GetPlayerSummaryResponse_200,
	})

	rec, c := testdata.InitEcho("/api/steam/user/search?q=testuser", "")

	if assert.NoError(t, GetSteamUserSearch(c)) {
		resBody, err := httputil.ReadBody[GetSteamUserSearchResponse](rec.Result())
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, nil, err)
		assert.Equal(t, SuccessResponse_200, resBody)
	}
}

func TestGetUser_MatchVanityURL(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerSummaryPath,
		Response: GetPlayerSummaryResponse_404,
		Querys:   map[string]string{"steamids": "testuser"},
	})
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerSummaryPath,
		Response: GetPlayerSummaryResponse_200,
		Querys:   map[string]string{"steamids": "0"},
	})
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.ResolveVanityURLPath,
		Response: ResolveVanityURLResponse_Match,
	})

	rec, c := testdata.InitEcho("/api/steam/user/search?q=testuser", "")

	if assert.NoError(t, GetSteamUserSearch(c)) {
		resBody, err := httputil.ReadBody[GetSteamUserSearchResponse](rec.Result())
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, nil, err)
		assert.Equal(t, SuccessResponse_200, resBody)
	}
}

func TestGetUser_NoMatchSteamID(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerSummaryPath,
		Response: GetPlayerSummaryResponse_404,
	})
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.ResolveVanityURLPath,
		Response: ResolveVanityURLResponse_Match,
	})

	rec, c := testdata.InitEcho("/api/steam/user/search?q=testuser", "")

	if assert.NoError(t, GetSteamUserSearch(c)) {
		resBody, err := httputil.ReadBody[GetSteamUserSearchResponse](rec.Result())
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, nil, err)
		assert.Equal(
			t,
			GetSteamUserSearchResponse{
				StatusCode: 200,
				Users:      make([]steamworks.Player, 0, 0),
			},
			resBody,
		)
	}
}

func TestGetUser_NoMatchVanityURL(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerSummaryPath,
		Response: GetPlayerSummaryResponse_404,
		Querys:   map[string]string{"steamids": "testuser"},
	})
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.GetPlayerSummaryPath,
		Response: GetPlayerSummaryResponse_200,
		Querys:   map[string]string{"steamids": "0"},
	})
	testdata.InitGock(testdata.GockConfig{
		Url:      "http://localhost:9999",
		Path:     steamworks.ResolveVanityURLPath,
		Response: ResolveVanityURLResponse_NoMatch,
	})

	rec, c := testdata.InitEcho("/api/steam/user/search?q=testuser", "")

	if assert.NoError(t, GetSteamUserSearch(c)) {
		resBody, err := httputil.ReadBody[GetSteamUserSearchResponse](rec.Result())
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, nil, err)
		assert.Equal(
			t,
			GetSteamUserSearchResponse{
				StatusCode: 200,
				Users:      make([]steamworks.Player, 0, 0),
			},
			resBody,
		)
	}
}

func TestGetUser_InvalidQ(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	rec, c := testdata.InitEcho("/api/steam/user/search", "")

	if assert.NoError(t, GetSteamUserSearch(c)) {
		resBody, err := httputil.ReadBody[httputil.Error](rec.Result())
		assert.Equal(t, 400, rec.Code)
		assert.Equal(t, nil, err)
		assert.Equal(
			t,
			httputil.Error{
				StatusCode: 400,
				ErrorCode:  "STEAM_USER_SEARCH_Q_NOT_FOUND",
				Message:    "q not found",
			},
			resBody,
		)
	}
}

func TestGetUser_ErrorSteam(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	rec, c := testdata.InitEcho("/api/steam/user/search?q=testuser", "")

	if assert.NoError(t, GetSteamUserSearch(c)) {
		resBody, err := httputil.ReadBody[httputil.Error](rec.Result())
		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, nil, err)
		assert.Equal(
			t,
			httputil.Error{
				StatusCode: 500,
				ErrorCode:  "STEAM_USER_SEARCH_ENVIROMENT_ERROR",
				Message:    "internal server error",
			},
			resBody,
		)
	}
}

var SuccessResponse_200 = GetSteamUserSearchResponse{
	StatusCode: 200,
	Users:      []steamworks.Player{steamworks.TestUser},
}

var GetPlayerSummaryResponse_200 = steamworks.GetPlayerSummaryResponse{
	Response: struct {
		Players []steamworks.Player "json:\"players\""
	}{
		Players: []steamworks.Player{steamworks.TestUser},
	},
}

var GetPlayerSummaryResponse_404 = steamworks.GetPlayerSummaryResponse{
	Response: struct {
		Players []steamworks.Player "json:\"players\""
	}{
		Players: make([]steamworks.Player, 0, 0),
	},
}

var ResolveVanityURLResponse_Match = steamworks.ResolveVanityURLResponse{
	Response: struct {
		Steamid string `json:"steamid"`
		Success int    `json:"success"`
		Message string `json:"message"`
	}{
		Steamid: steamworks.TestUser.SteamID,
		Success: 1,
		Message: "",
	},
}

var ResolveVanityURLResponse_NoMatch = steamworks.ResolveVanityURLResponse{
	Response: struct {
		Steamid string `json:"steamid"`
		Success int    `json:"success"`
		Message string `json:"message"`
	}{
		Success: 42,
		Message: "No match",
	},
}
