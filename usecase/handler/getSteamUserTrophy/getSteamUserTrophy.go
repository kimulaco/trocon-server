package GetSteamUserTrophyAPI

import (
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/kimulaco/trocon-server/interface/steamworks"
	"github.com/kimulaco/trocon-server/pkg/httputil"
	"github.com/kimulaco/trocon-server/pkg/stringsx"
	"github.com/labstack/echo/v4"
)

type Trophy struct {
	Success      bool                     `json:"success"`
	AppId        int                      `json:"appId"`
	GameName     string                   `json:"gameName"`
	Achievements []steamworks.Achievement `json:"trophies"`
}

type SuccessResponse struct {
	StatusCode int      `json:"statusCode"`
	Trophies   []Trophy `json:"trophies"`
}

func GetSteamUserTrophy(c echo.Context) error {
	s := steamworks.NewSteamworks()
	if s.InvalidEnv() {
		return c.JSON(httputil.NewError500("STEAM_USER_TROPHY_ENVIROMENT_ERROR", ""))
	}

	steamid := c.Param("steamid")
	if steamid == "" {
		return c.JSON(httputil.NewError400("STEAM_USER_TROPHY_STEAMID_NOT_FOUND", "steamid not found"))
	}

	appid := c.QueryParam("appid")
	appids := strings.Split(appid, ",")
	appidLen := len(appids)
	if appid == "" || appidLen <= 0 {
		return c.JSON(httputil.NewError400("STEAM_USER_TROPHY_APPID_NOT_FOUND", "appid not found"))
	}

	trophiesChan := make(chan Trophy, appidLen)
	trophies := make([]Trophy, 0, appidLen)

	for i := 0; i < appidLen; i++ {
		go getPlayerArchivementWithCh(s, steamid, appids[i], trophiesChan)
	}

	for i := 0; i < appidLen; i++ {
		trophy := <-trophiesChan
		trophies = append(trophies, trophy)
	}

	trophies = sortTrophy(trophies, appids)

	return c.JSON(http.StatusOK, SuccessResponse{
		StatusCode: http.StatusOK,
		Trophies:   trophies,
	})
}

func getPlayerArchivementWithCh(
	s steamworks.Steamworks, steamid string, appid string, ch chan Trophy,
) {
	appidInt, _ := strconv.Atoi(appid)
	game, err := s.GetPlayerAchievements(steamid, appid)
	if err != nil {
		log.Print("STEAM_USER_TROPHY_INTERNAL_ERROR: appid:" + appid + " " + err.Error())
		game.Success = false
	}

	ch <- Trophy{
		Success:      GetSuccess(game),
		AppId:        appidInt,
		GameName:     game.GameName,
		Achievements: game.Achievements,
	}
}

func sortTrophy(trophies []Trophy, appids []string) []Trophy {
	sort.SliceStable(trophies, func(i, j int) bool {
		trophyAIndex := stringsx.IndexOfString(strconv.Itoa(trophies[i].AppId), appids)
		trophyBIndex := stringsx.IndexOfString(strconv.Itoa(trophies[j].AppId), appids)

		return trophyAIndex < trophyBIndex
	})

	return trophies
}
