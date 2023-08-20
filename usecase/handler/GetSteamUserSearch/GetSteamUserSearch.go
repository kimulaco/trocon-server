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
	StatusCode int                 `json:"statusCode"`
	Users      []steamworks.Player `json:"users"`
}

// GetSteamUserSearch
//
//	@Tags			Steam User
//	@Description	Get steam user by Steam ID or Steam vanity URL.
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

	MAX_LENGTH := 2
	playerChan := make(chan steamworks.Player, MAX_LENGTH)
	users := make([]steamworks.Player, 0, MAX_LENGTH)

	go getPlayerWithID(s, q, playerChan)
	go getPlayerWithVanityURL(s, q, playerChan)

	for i := 0; i < MAX_LENGTH; i++ {
		_user := <-playerChan
		if !_user.IsEmpty() {
			users = append(users, _user)
		}
	}

	return c.JSON(http.StatusOK, GetSteamUserSearchResponse{
		StatusCode: http.StatusOK,
		Users:      users,
	})
}

func getPlayerWithID(
	s steamworks.Steamworks, steamid string, ch chan steamworks.Player,
) {
	player, err := s.GetPlayerSummary(steamid)
	if err != nil {
		log.Println(err.Error())
	}

	ch <- player
}

func getPlayerWithVanityURL(
	s steamworks.Steamworks, q string, ch chan steamworks.Player,
) {
	steamid, err := s.ResolveVanityURL(q)
	if err != nil {
		log.Println(err.Error())
		ch <- steamworks.Player{}
		return
	}

	player, err := s.GetPlayerSummary(steamid)
	if err != nil {
		log.Println(err.Error())
	}

	ch <- player
}

func errorLog(s string) {
	log.Print(s)
	sentry.CaptureException(errors.New(s))
}
