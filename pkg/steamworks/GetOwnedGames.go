package steamworks

import (
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

func (s Steamworks) GetOwnedGames(steamid string) ([]OwnedGame, error) {
	var ownedGames []OwnedGame

	apiUrl, _ := urlx.NewUrlx(s.APIBaseURL + "/IPlayerService/GetOwnedGames/v1/")
	apiUrl.AddQuery("key", s.APIKey)
	apiUrl.AddQuery("steamid", steamid)
	apiUrl.AddQuery("include_appinfo", "true")
	apiUrl.AddQuery("include_played_free_games", "true")

	res, err := http.Get(apiUrl.ToString())
	if err != nil {
		return ownedGames, err
	}
	defer res.Body.Close()

	resBody, err := httputil.ReadBody[GetOwnedGamesResponse](res)
	if err != nil {
		return ownedGames, err
	}

	ownedGames = resBody.Response.Games

	return ownedGames, nil
}
