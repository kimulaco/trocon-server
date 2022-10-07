package getSteamUserAPI

import (
	"log"
	"net/http"

	"github.com/kimulaco/g-trophy/pkg/httputil"
	"github.com/kimulaco/g-trophy/pkg/steamworks"
	"github.com/labstack/echo/v4"
)

type GetUserResponse struct {
	StatusCode int                    `json:"statusCode"`
	User       steamworks.Player      `json:"user"`
	Games      []steamworks.OwnedGame `json:"games"`
}

func GetUser(c echo.Context) error {
	steamid := c.Param("steamid")

	steamApiKey := steamworks.NewApiKey()
	if !steamApiKey.HasKey() {
		log.Print("STEAM_USER_ENVIROMENT_ERROR: undefined steam API key")
		return c.JSON(httputil.NewError500("STEAM_USER_ENVIROMENT_ERROR", ""))
	}

	player, err := steamworks.GetPlayerSummary(steamApiKey.Key, steamid)
	if err != nil {
		log.Print("STEAM_USER_INTERNAL_ERROR: " + err.Error())
		return c.JSON(httputil.NewError500("STEAM_USER_INTERNAL_ERROR", ""))
	}

	if player.IsEmpty() {
		return c.JSON(httputil.NewError404("STEAM_USER_NOT_FOUND", "user not found"))
	}

	games, err := steamworks.GetOwnedGames(steamApiKey.Key, steamid)
	if err != nil {
		log.Print("STEAM_OWNED_GAME_INTERNAL_ERROR: " + err.Error())
		return c.JSON(httputil.NewError500("STEAM_OWNED_GAME_INTERNAL_ERROR", ""))
	}

	return c.JSON(http.StatusOK, GetUserResponse{
		StatusCode: http.StatusOK,
		User:       player,
		Games:      games,
	})
}
