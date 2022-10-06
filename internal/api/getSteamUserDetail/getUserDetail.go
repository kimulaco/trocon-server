package GetSteamUserDetailAPI

import (
	"net/http"

	"github.com/kimulaco/g-trophy/pkg/httputil"
	"github.com/kimulaco/g-trophy/pkg/steamworks"
	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	StatusCode int               `json:"statusCode"`
	User       steamworks.Player `json:"user"`
}

func GetUserDetail(c echo.Context) error {
	steamApiKey := steamworks.NewApiKey()
	steamid := c.Param("steamid")

	if !steamApiKey.HasKey() {
		return c.JSON(httputil.NewError500("STEAM_USER_ENVIROMENT_ERROR", ""))
	}

	player, _ := steamworks.GetPlayerSummary(steamApiKey.Key, steamid)

	if player.IsEmpty() {
		return c.JSON(httputil.NewError404("STEAM_USER_NOT_FOUND", "user not found"))
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		StatusCode: http.StatusOK,
		User:       player,
	})
}
