package steamworks

import (
	"errors"
	"net/http"

	"github.com/kimulaco/trophy-comp-server/pkg/httputil"
	"github.com/kimulaco/trophy-comp-server/pkg/urlx"
)

type GetOwnedGamesResponse struct {
	Response struct {
		GameCount int         `json:"game_count"`
		Games     []OwnedGame `json:"games"`
	} `json:"response"`
}

const GetOwnedGamesPath = "/IPlayerService/GetOwnedGames/v1/"

func (s Steamworks) GetOwnedGames(steamid string) ([]OwnedGame, error) {
	ownedGames := make([]OwnedGame, 0)

	if steamid == "" {
		return ownedGames, errors.New("steamid is required")
	}

	apiUrl, _ := urlx.NewUrlx(s.APIBaseURL + GetOwnedGamesPath)
	apiUrl.AddQuery("key", s.APIKey)
	apiUrl.AddQuery("steamid", steamid)
	apiUrl.AddQuery("include_appinfo", "true")
	apiUrl.AddQuery("include_played_free_games", "true")

	res, err := http.Get(apiUrl.ToString())
	if err != nil {
		return ownedGames, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ownedGames, errors.New(res.Status)
	}

	resBody, err := httputil.ReadBody[GetOwnedGamesResponse](res)
	if err != nil {
		return ownedGames, err
	}

	ownedGames = resBody.Response.Games

	return ownedGames, nil
}
