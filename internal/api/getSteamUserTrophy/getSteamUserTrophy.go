package GetSteamUserTrophyAPI

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/kimulaco/trophy-comp-server/pkg/httputil"
	"github.com/kimulaco/trophy-comp-server/pkg/steamworks"
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
	steamid := c.Param("steamid")
	appids := strings.Split(c.QueryParam("appid"), ",")
	s := steamworks.NewSteamworks()

	if s.InvalidEnv() {
		log.Print("STEAM_USER_TROPHY_ENVIROMENT_ERROR: undefined steam API key")
		return c.JSON(httputil.NewError500("STEAM_USER_TROPHY_ENVIROMENT_ERROR", ""))
	}

	var trophies []Trophy

	for _, appid := range appids {
		appidInt, _ := strconv.Atoi(appid)
		game, err := s.GetPlayerAchievements(steamid, appid)
		success := true

		if err != nil {
			log.Print("STEAM_USER_TROPHY_INTERNAL_ERROR: " + err.Error())
			success = false
		}

		trophies = append(trophies, Trophy{
			Success:      success,
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
