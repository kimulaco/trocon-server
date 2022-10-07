package main

import (
	"github.com/labstack/echo/v4"

	GetSteamOwnedGamesAPI "github.com/kimulaco/g-trophy/internal/api/getSteamOwnedGames"
	GetSteamUserDetailAPI "github.com/kimulaco/g-trophy/internal/api/getSteamUserDetail"
	GetSteamUserTrophyAPI "github.com/kimulaco/g-trophy/internal/api/getSteamUserTrophy"
)

func main() {
	e := echo.New()

	e.GET("/api/steam/user/:steamid", GetSteamUserDetailAPI.GetUserDetail)
	e.GET("/api/steam/user/:steamid/games", GetSteamOwnedGamesAPI.GetOwnedGames)
	e.GET("/api/steam/user/:steamid/trophy", GetSteamUserTrophyAPI.GetSteamUserTrophy)

	e.Logger.Fatal(e.Start(":9000"))
}
