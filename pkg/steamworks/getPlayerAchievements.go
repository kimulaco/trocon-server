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

type GetPlayerAchievementsResponse struct {
	PlayerStats GetPlayerAchievementsResponseOwnedGame `json:"playerstats"`
}

const GetPlayerAchievementsPath = "/ISteamUserStats/GetPlayerAchievements/v1/"

func (s Steamworks) GetPlayerAchievements(
	steamid string,
	appid string,
) (GetPlayerAchievementsResponseOwnedGame, error) {
	var response GetPlayerAchievementsResponseOwnedGame

	apiUrl, _ := urlx.NewUrlx(s.APIBaseURL + GetPlayerAchievementsPath)
	apiUrl.AddQuery("key", s.APIKey)
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

	if response.GameName == "" {
		return response, errors.New("appid:" + appid + " not found")
	}

	return response, nil
}
