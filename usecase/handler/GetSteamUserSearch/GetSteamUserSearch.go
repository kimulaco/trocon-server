package GetSteamUserSearchAPI

import (
	"errors"
	"log"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/kimulaco/trocon-server/interface/steamworks"
	"github.com/kimulaco/trocon-server/pkg/httputil"
	"github.com/labstack/echo/v4"
)

type GetSteamUserSearchResponse struct {
	StatusCode int    `json:"statusCode"`
	Steamid    string `json:"steamid"`
}

// GetSteamUserSearch
//
//	@Tags			Steam User
//	@Description	Get steamid searched by keyword.
//	@Accept			json
//	@Produce		json
//	@Param			q	query		string	true	"Steam ID or Steam vanity URL"
//	@Success		200	{object}	GetSteamUserSearchResponse
//	@Failure		400	{object}	httputil.Error
//	@Failure		404	{object}	httputil.Error
//	@Failure		500	{object}	httputil.Error
//	@Router			/api/steam/user/search [get]
func GetSteamUserSearch(c echo.Context) error {
	s := steamworks.NewSteamworks()
	if s.InvalidEnv() {
		errorLog("STEAM_USER_SEARCH_ENVIROMENT_ERROR: undefined steam API key")
		return c.JSON(httputil.NewError500("STEAM_USER_SEARCH_ENVIROMENT_ERROR", ""))
	}

	q := c.QueryParam("q")
	if q == "" {
		return c.JSON(httputil.NewError400("STEAM_USER_SEARCH_Q_NOT_FOUND", "q not found"))
	}

	queryPlayer, err := s.GetPlayerSummary(q)
	if err == nil && !queryPlayer.IsEmpty() {
		return c.JSON(http.StatusOK, GetSteamUserSearchResponse{
			StatusCode: http.StatusOK,
			Steamid:    queryPlayer.SteamID,
		})
	}
	if err != nil {
		log.Println(err.Error())
	}

	steamid, err := s.ResolveVanityURL(q)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(httputil.NewError404(
			"STEAM_USER_SEARCH_FAIL_RESOLVE_VANITY_URL",
			"User not found",
		))
	}

	idPlayer, err := s.GetPlayerSummary(steamid)
	if err == nil && !idPlayer.IsEmpty() {
		return c.JSON(http.StatusOK, GetSteamUserSearchResponse{
			StatusCode: http.StatusOK,
			Steamid:    idPlayer.SteamID,
		})
	}
	if err != nil {
		log.Println(err.Error())
	}

	return c.JSON(httputil.NewError404(
		"STEAM_USER_SEARCH_FAIL_GET_PLAYER_SUMMARY",
		"User not found",
	))
}

func errorLog(s string) {
	log.Print(s)
	sentry.CaptureException(errors.New(s))
}
