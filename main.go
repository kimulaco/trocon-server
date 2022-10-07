package main

import (
	"github.com/labstack/echo/v4"

	GetSteamUserAPI "github.com/kimulaco/g-trophy/internal/api/getSteamUser"
	GetSteamUserDetailAPI "github.com/kimulaco/g-trophy/internal/api/getSteamUserDetail"
	GetSteamUserGamesAPI "github.com/kimulaco/g-trophy/internal/api/getSteamUserGames"
	GetSteamUserTrophyAPI "github.com/kimulaco/g-trophy/internal/api/getSteamUserTrophy"
)

func main() {
	e := echo.New()

	e.GET("/api/steam/user/:steamid", GetSteamUserAPI.GetUser)
	e.GET("/api/steam/user/:steamid/detail", GetSteamUserDetailAPI.GetUserDetail)
	e.GET("/api/steam/user/:steamid/games", GetSteamUserGamesAPI.GetSteamUserGames)
	e.GET("/api/steam/user/:steamid/trophy", GetSteamUserTrophyAPI.GetSteamUserTrophy)

	e.Logger.Fatal(e.Start(":9000"))
}
