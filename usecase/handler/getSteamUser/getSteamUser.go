package GetSteamUserAPI

import (
	"errors"
	"log"
	"net/http"

	"github.com/getsentry/sentry-go"
	Game "github.com/kimulaco/trocon-server/domain/game"
	"github.com/kimulaco/trocon-server/interface/steamworks"
	"github.com/kimulaco/trocon-server/pkg/httputil"
	"github.com/labstack/echo/v4"
)

type GetSteamUserResponse struct {
	StatusCode int               `json:"statusCode"`
	User       steamworks.Player `json:"user"`
	Games      []Game.Game       `json:"games"`
}

// GetSteamUser
//
//	@Tags			Steam User
//	@Description	Get server status.
//	@Accept			json
//	@Produce		json
//	@Param			steamid	path		string	true	"Steam ID"
//	@Success		200		{object}	GetSteamUserResponse
//	@Failure		400		{object}	httputil.Error
//	@Failure		404		{object}	httputil.Error
//	@Failure		500		{object}	httputil.Error
//	@Router			/api/steam/user/:steamid [get]
func GetSteamUser(c echo.Context) error {
	s := steamworks.NewSteamworks()
	if s.InvalidEnv() {
		errorLog("STEAM_USER_ENVIROMENT_ERROR: undefined steam API key")
		return c.JSON(httputil.NewError500("STEAM_USER_ENVIROMENT_ERROR", ""))
	}

	steamid := c.Param("steamid")
	if steamid == "" {
		return c.JSON(httputil.NewError400("STEAM_USER_STEAMID_NOT_FOUND", "steamid not found"))
	}

	var playerCh UserChannel
	var ownedGamesCh UserChannel
	ch := make(chan UserChannel, 2)

	go getPlayerSummary(s, steamid, ch)
	go getOwnedGames(s, steamid, ch)

	// TODO: Refactor channel receiving
	userCh1, userCh2 := <-ch, <-ch

	for _, c := range []UserChannel{userCh1, userCh2} {
		if len(c.player.SteamID) > 0 {
			playerCh = c
		} else {
			ownedGamesCh = c
		}
	}

	if playerCh.err != nil {
		errorLog("STEAM_USER_INTERNAL_ERROR: " + playerCh.err.Error())
		return c.JSON(httputil.NewError500("STEAM_USER_INTERNAL_ERROR", ""))
	}

	if playerCh.player.IsEmpty() {
		return c.JSON(httputil.NewError404("STEAM_USER_NOT_FOUND", "user not found"))
	}

	if ownedGamesCh.err != nil {
		errorLog("STEAM_OWNED_GAME_INTERNAL_ERROR: " + ownedGamesCh.err.Error())
		return c.JSON(httputil.NewError500("STEAM_OWNED_GAME_INTERNAL_ERROR", ""))
	}

	return c.JSON(http.StatusOK, GetSteamUserResponse{
		StatusCode: http.StatusOK,
		User:       playerCh.player,
		Games:      steamworks.MapOwnedGamesToGames(ownedGamesCh.ownedGames),
	})
}

func errorLog(s string) {
	log.Print(s)
	sentry.CaptureException(errors.New(s))
}

type UserChannel struct {
	player     steamworks.Player
	ownedGames []steamworks.OwnedGame
	err        error
}

func getPlayerSummary(s steamworks.Steamworks, steamid string, ch chan UserChannel) {
	player, err := s.GetPlayerSummary(steamid)

	ch <- UserChannel{
		player: player,
		err:    err,
	}
}

func getOwnedGames(s steamworks.Steamworks, steamid string, ch chan UserChannel) {
	ownedGames, err := s.GetOwnedGames(steamid)

	ch <- UserChannel{
		ownedGames: ownedGames,
		err:        err,
	}
}
