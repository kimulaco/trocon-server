package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	GetStatusAPI "github.com/kimulaco/trophy-comp-server/internal/api/getStatus"
	GetSteamUserAPI "github.com/kimulaco/trophy-comp-server/internal/api/getSteamUser"
	GetSteamUserTrophyAPI "github.com/kimulaco/trophy-comp-server/internal/api/getSteamUserTrophy"
	"github.com/kimulaco/trophy-comp-server/internal/config"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.GetCORSAllowOrigins(),
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/api/status", GetStatusAPI.GetStatus)

	e.GET("/api/steam/user/:steamid", GetSteamUserAPI.GetUser)
	e.GET("/api/steam/user/:steamid/trophy", GetSteamUserTrophyAPI.GetSteamUserTrophy)

	e.Logger.Fatal(e.Start(config.GetListenPort("9000")))
}
