package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	GetSteamUserAPI "github.com/kimulaco/g-trophy-server/internal/api/getSteamUser"
	GetSteamUserDetailAPI "github.com/kimulaco/g-trophy-server/internal/api/getSteamUserDetail"
	GetSteamUserGamesAPI "github.com/kimulaco/g-trophy-server/internal/api/getSteamUserGames"
	GetSteamUserTrophyAPI "github.com/kimulaco/g-trophy-server/internal/api/getSteamUserTrophy"
	"github.com/kimulaco/g-trophy-server/internal/config"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.GetCORSAllowOrigins(),
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/api/steam/user/:steamid", GetSteamUserAPI.GetUser)
	e.GET("/api/steam/user/:steamid/detail", GetSteamUserDetailAPI.GetUserDetail)
	e.GET("/api/steam/user/:steamid/games", GetSteamUserGamesAPI.GetSteamUserGames)
	e.GET("/api/steam/user/:steamid/trophy", GetSteamUserTrophyAPI.GetSteamUserTrophy)

	e.Logger.Fatal(e.Start(":9000"))
}
