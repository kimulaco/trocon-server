package GetSteamUserTrophyAPI

import (
	"log"
	"net/http"

	"github.com/kimulaco/g-trophy-server/pkg/httputil"
	"github.com/kimulaco/g-trophy-server/pkg/steamworks"
	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	StatusCode int                      `json:"statusCode"`
	GameName   string                   `json:"gameName"`
	Trophies   []steamworks.Achievement `json:"trophies"`
}

func GetSteamUserTrophy(c echo.Context) error {
	steamid := c.Param("steamid")
	appid := c.QueryParam("appid")

	steamApiKey := steamworks.NewApiKey()
	if !steamApiKey.HasKey() {
		log.Print("STEAM_USER_TROPHY_ENVIROMENT_ERROR: undefined steam API key")
		return c.JSON(httputil.NewError500("STEAM_USER_TROPHY_ENVIROMENT_ERROR", ""))
	}

	trophies, err := steamworks.GetPlayerAchievements(steamApiKey.Key, steamid, appid)
	if err != nil {
		log.Print("STEAM_USER_TROPHY_INTERNAL_ERROR: " + err.Error())
		return c.JSON(httputil.NewError500("STEAM_USER_TROPHY_INTERNAL_ERROR", ""))
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		StatusCode: http.StatusOK,
		GameName:   trophies.GameName,
		Trophies:   trophies.Achievements,
	})
}
