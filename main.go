package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoswagger "github.com/swaggo/echo-swagger"

	_ "github.com/kimulaco/trocon-server/docs"
	"github.com/kimulaco/trocon-server/interface/sentry"
	APIConfig "github.com/kimulaco/trocon-server/usecase/config"
	GetSteamUserSearchAPI "github.com/kimulaco/trocon-server/usecase/handler/GetSteamUserSearch"
	GetStatusAPI "github.com/kimulaco/trocon-server/usecase/handler/getStatus"
	GetSteamUserAPI "github.com/kimulaco/trocon-server/usecase/handler/getSteamUser"
	GetSteamUserTrophyAPI "github.com/kimulaco/trocon-server/usecase/handler/getSteamUserTrophy"
)

//	@title			Trocon API
//	@version		0.1
//	@description	Rest API for Trocon.
func main() {
	env := os.Getenv("BUILD_MODE")
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

	e.GET("/api/steam/user/:steamid", GetSteamUserAPI.GetSteamUser)
	e.GET("/api/steam/user/:steamid/trophy", GetSteamUserTrophyAPI.GetSteamUserTrophy)
	e.GET("/api/steam/user/search", GetSteamUserSearchAPI.GetSteamUserSearch)

	if env != "production" {
		e.GET("/swagger/*", echoswagger.WrapHandler)
	}

	e.Logger.Fatal(e.Start(APIConfig.GetListenPort("9000")))
}
