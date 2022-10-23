package GetSteamUserAPI

import (
	"log"
	"net/http"

	"github.com/kimulaco/trophy-comp-server/pkg/httputil"
	"github.com/kimulaco/trophy-comp-server/pkg/steamworks"
	"github.com/labstack/echo/v4"
)

type GetUserResponse struct {
	StatusCode int               `json:"statusCode"`
	User       steamworks.Player `json:"user"`
	Games      []steamworks.Game `json:"games"`
}

func GetUser(c echo.Context) error {
	s := steamworks.NewSteamworks()
	if s.InvalidEnv() {
		log.Print("STEAM_USER_ENVIROMENT_ERROR: undefined steam API key")
		return c.JSON(httputil.NewError500("STEAM_USER_ENVIROMENT_ERROR", ""))
	}

	steamid := c.Param("steamid")
	if steamid == "" {
		return c.JSON(httputil.NewError400("STEAM_USER_STEAMID_NOT_FOUND", "steamid not found"))
	}

	player, err := s.GetPlayerSummary(steamid)
	if err != nil {
		log.Print("STEAM_USER_INTERNAL_ERROR: " + err.Error())
		return c.JSON(httputil.NewError500("STEAM_USER_INTERNAL_ERROR", ""))
	}

	if player.IsEmpty() {
		return c.JSON(httputil.NewError404("STEAM_USER_NOT_FOUND", "user not found"))
	}

	ownedGames, err := s.GetOwnedGames(steamid)
	if err != nil {
		log.Print("STEAM_OWNED_GAME_INTERNAL_ERROR: " + err.Error())
		return c.JSON(httputil.NewError500("STEAM_OWNED_GAME_INTERNAL_ERROR", ""))
	}

	return c.JSON(http.StatusOK, GetUserResponse{
		StatusCode: http.StatusOK,
		User:       player,
		Games:      steamworks.MapOwnedGamesToGames(ownedGames),
	})
}
