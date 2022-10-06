package main

import (
	"github.com/labstack/echo/v4"

	GetSteamUserDetailAPI "github.com/kimulaco/g-trophy/internal/api/getSteamUserDetail"
)

func main() {
	e := echo.New()

	e.GET("/api/steam/user/:steamid", GetSteamUserDetailAPI.GetUserDetail)

	e.Logger.Fatal(e.Start(":9000"))
}
