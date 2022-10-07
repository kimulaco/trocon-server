package GetOwnedGamesAPI

import (
	"log"
	"net/http"

	"github.com/kimulaco/g-trophy/pkg/httputil"
	"github.com/kimulaco/g-trophy/pkg/steamworks"
	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	StatusCode int                    `json:"statusCode"`
	Games      []steamworks.OwnedGame `json:"games"`
}

func GetOwnedGames(c echo.Context) error {
	steamid := c.Param("steamid")

	steamApiKey := steamworks.NewApiKey()
	if !steamApiKey.HasKey() {
		log.Print("STEAM_OWNED_GAME_ENVIROMENT_ERROR: undefined steam API key")
		return c.JSON(httputil.NewError500("STEAM_OWNED_GAME_ENVIROMENT_ERROR", ""))
	}

	games, err := steamworks.GetOwnedGames(steamApiKey.Key, steamid)
	if err != nil {
		log.Print("STEAM_OWNED_GAME_INTERNAL_ERROR: " + err.Error())
		return c.JSON(httputil.NewError500("STEAM_OWNED_GAME_INTERNAL_ERROR", ""))
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		StatusCode: http.StatusOK,
		Games:      games,
	})
}
