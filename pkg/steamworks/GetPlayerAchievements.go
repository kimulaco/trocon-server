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

	if steamid == "" {
		return response, errors.New("steamid is required")
	}

	if appid == "" {
		return response, errors.New("appid is required")
	}

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

	if res.StatusCode != http.StatusOK {
		return response, errors.New(res.Status)
	}

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
