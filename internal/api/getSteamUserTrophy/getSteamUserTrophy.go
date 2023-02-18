package GetSteamUserTrophyAPI

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/kimulaco/trocon-server/pkg/httputil"
	"github.com/kimulaco/trocon-server/pkg/steamworks"
	"github.com/labstack/echo/v4"
)

type Trophy struct {
	Success      bool                     `json:"success"`
	AppId        int                      `json:"appId"`
	GameName     string                   `json:"gameName"`
	Achievements []steamworks.Achievement `json:"trophies"`
}

type SuccessResponse struct {
	StatusCode int      `json:"statusCode"`
	Trophies   []Trophy `json:"trophies"`
}

func GetSteamUserTrophy(c echo.Context) error {
	s := steamworks.NewSteamworks()
	if s.InvalidEnv() {
		return c.JSON(httputil.NewError500("STEAM_USER_TROPHY_ENVIROMENT_ERROR", ""))
	}

	steamid := c.Param("steamid")
	if steamid == "" {
		return c.JSON(httputil.NewError400("STEAM_USER_TROPHY_STEAMID_NOT_FOUND", "steamid not found"))
	}

	appid := c.QueryParam("appid")
	appids := strings.Split(appid, ",")
	if appid == "" || len(appids) <= 0 {
		return c.JSON(httputil.NewError400("STEAM_USER_TROPHY_APPID_NOT_FOUND", "appid not found"))
	}

	var trophies []Trophy

	for _, appid := range appids {
		appidInt, _ := strconv.Atoi(appid)
		game, err := s.GetPlayerAchievements(steamid, appid)

		if err != nil {
			log.Print("STEAM_USER_TROPHY_INTERNAL_ERROR: appid:" + appid + " " + err.Error())
			game.Success = false
		}

		trophies = append(trophies, Trophy{
			Success:      game.Success,
			AppId:        appidInt,
			GameName:     game.GameName,
			Achievements: game.Achievements,
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		StatusCode: http.StatusOK,
		Trophies:   trophies,
	})
}
