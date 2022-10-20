package steamworks

import (
	"errors"
	"net/http"

	"github.com/kimulaco/trophy-comp-server/pkg/httputil"
	"github.com/kimulaco/trophy-comp-server/pkg/urlx"
)

type GetPlayerAchievementsResponseOwnedGame struct {
	GameName     string        `json:"gameName"`
	Achievements []Achievement `json:"achievements"`
}

func (g GetPlayerAchievementsResponseOwnedGame) IsEmpty() bool {
	return g.GameName == ""
}

type GetPlayerAchievementsResponse struct {
	PlayerStats GetPlayerAchievementsResponseOwnedGame `json:"playerstats"`
}

func GetPlayerAchievements(
	key string,
	steamid string,
	appid string,
) (GetPlayerAchievementsResponseOwnedGame, error) {
	const _API_URL = "https://api.steampowered.com/ISteamUserStats/GetPlayerAchievements/v1/"
	var response GetPlayerAchievementsResponseOwnedGame

	apiUrl, _ := urlx.NewUrlx(_API_URL)
	apiUrl.AddQuery("key", key)
	apiUrl.AddQuery("steamid", steamid)
	apiUrl.AddQuery("appid", appid)
	apiUrl.AddQuery("l", "japanese")

	res, err := http.Get(apiUrl.ToString())
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	resBody, err := httputil.ReadBody[GetPlayerAchievementsResponse](res)
	if err != nil {
		return response, err
	}

	response = resBody.PlayerStats

	if response.Achievements == nil {
		response.Achievements = []Achievement{}
	}

	if response.IsEmpty() {
		return response, errors.New("appid:" + appid + " not found")
	}

	return response, nil
}
