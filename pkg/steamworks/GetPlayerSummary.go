package steamworks

import (
	"errors"
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

	if steamid == "" {
		return player, errors.New("steamid is required")
	}

	apiUrl, _ := urlx.NewUrlx(s.APIBaseURL + GetPlayerSummaryPath)
	apiUrl.AddQuery("key", s.APIKey)
	apiUrl.AddQuery("steamids", steamid)

	res, err := http.Get(apiUrl.ToString())
	if err != nil {
		return player, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return player, errors.New(res.Status)
	}

	resBody, err := httputil.ReadBody[GetPlayerSummaryResponse](res)
	if err != nil {
		return player, err
	}

	if len(resBody.Response.Players) > 0 {
		player = resBody.Response.Players[0]
	}

	return player, nil
}
