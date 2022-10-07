package steamworks

import (
	"net/http"

	"github.com/kimulaco/g-trophy/pkg/httputil"
	"github.com/kimulaco/g-trophy/pkg/urlx"
)

type GetPlayerSummaryResponsePlayers struct {
	Players []Player `json:"players"`
}

type GetPlayerSummaryResponse struct {
	Response GetPlayerSummaryResponsePlayers `json:"response"`
}

func GetPlayerSummary(key string, steamid string) (Player, error) {
	const _API_URL = "https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/"
	var player Player

	apiUrl, _ := urlx.NewUrlx(_API_URL)
	apiUrl.AddQuery("key", key)
	apiUrl.AddQuery("steamids", steamid)

	res, err := http.Get(apiUrl.ToString())
	if err != nil {
		return player, err
	}
	defer res.Body.Close()

	resBody, err := httputil.ReadBody[GetPlayerSummaryResponse](res)
	if err != nil {
		return player, err
	}

	if len(resBody.Response.Players) > 0 {
		player = resBody.Response.Players[0]
	}

	return player, nil
}
