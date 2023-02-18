package steamworks

import (
	"errors"
	"net/http"

	"github.com/kimulaco/trocon-server/pkg/httputil"
	"github.com/kimulaco/trocon-server/pkg/urlx"
)

type GetPlayerAchievementsResponseOwnedGame struct {
	GameName     string        `json:"gameName"`
	Achievements []Achievement `json:"achievements"`
	Error string `json:"error"`
	Success bool `json:"success"`
}

type GetPlayerAchievementsResponse struct {
	PlayerStats GetPlayerAchievementsResponseOwnedGame `json:"playerstats"`
}

const GetPlayerAchievementsPath = "/ISteamUserStats/GetPlayerAchievements/v1/"

func (s Steamworks) GetPlayerAchievements(
	steamid string,
	appid string,
) (GetPlayerAchievementsResponseOwnedGame, error) {
	response := GetPlayerAchievementsResponseOwnedGame{
		GameName: "",
		Achievements: make([]Achievement, 0),
		Error: "",
		Success: false,
	}

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

	resBody, err := httputil.ReadBody[GetPlayerAchievementsResponse](res)
	if err != nil {
		if res.StatusCode != http.StatusOK {
			return response, errors.New(res.Status)
		}
		return response, err
	}

	response = resBody.PlayerStats

	if response.Achievements == nil || len(response.Achievements) <= 0 {
		response.Achievements = make([]Achievement, 0)
	}

	if !response.Success || res.StatusCode != http.StatusOK {
		response.Success = false
	}

	return response, nil
}
