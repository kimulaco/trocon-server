package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/kimulaco/trocon-server/interface/sentry"
	APIConfig "github.com/kimulaco/trocon-server/usecase/config"
	GetSteamUserSearchAPI "github.com/kimulaco/trocon-server/usecase/handler/GetSteamUserSearch"
	GetStatusAPI "github.com/kimulaco/trocon-server/usecase/handler/getStatus"
	GetSteamUserAPI "github.com/kimulaco/trocon-server/usecase/handler/getSteamUser"
	GetSteamUserTrophyAPI "github.com/kimulaco/trocon-server/usecase/handler/getSteamUserTrophy"
)

func main() {
	err := sentry.Init()
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	defer func() {
		sentry.Recover()
	}()

	e := echo.New()

	e.Use(sentry.NewMiddleware())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: APIConfig.GetCORSAllowOrigins(),
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/api/status", GetStatusAPI.GetStatus)

	e.GET("/api/steam/user/:steamid", GetSteamUserAPI.GetUser)
	e.GET("/api/steam/user/:steamid/trophy", GetSteamUserTrophyAPI.GetSteamUserTrophy)
	e.GET("/api/steam/user/search", GetSteamUserSearchAPI.GetSteamUserSearch)

	e.Logger.Fatal(e.Start(APIConfig.GetListenPort("9000")))
}
