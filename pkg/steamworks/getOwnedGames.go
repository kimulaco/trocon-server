package steamworks

import (
	"net/http"

	"github.com/kimulaco/trophy-comp-server/pkg/httputil"
	"github.com/kimulaco/trophy-comp-server/pkg/urlx"
)

type GetOwnedGamesResponseOwnedGame struct {
	GameCount int         `json:"game_count"`
	Games     []OwnedGame `json:"games"`
}

type GetOwnedGamesResponse struct {
	Response GetOwnedGamesResponseOwnedGame `json:"response"`
}

func GetOwnedGames(key string, steamid string) ([]OwnedGame, error) {
	const _API_URL = "https://api.steampowered.com/IPlayerService/GetOwnedGames/v1/"
	var ownedGames []OwnedGame

	apiUrl, _ := urlx.NewUrlx(_API_URL)
	apiUrl.AddQuery("key", key)
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
