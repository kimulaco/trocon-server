package steamworks

import (
	"net/http"

	"github.com/kimulaco/trophy-comp-server/pkg/httputil"
	"github.com/kimulaco/trophy-comp-server/pkg/urlx"
)

type GetPlayerSummaryResponse struct {
	Response struct {
		Players []Player `json:"players"`
	} `json:"response"`
}

const GetPlayerSummaryPath = "/ISteamUser/GetPlayerSummaries/v2/"

func (s Steamworks) GetPlayerSummary(steamid string) (Player, error) {
	var player Player

	apiUrl, _ := urlx.NewUrlx(s.APIBaseURL + GetPlayerSummaryPath)
	apiUrl.AddQuery("key", s.APIKey)
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
