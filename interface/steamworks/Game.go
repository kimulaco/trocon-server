package steamworks

import (
	"strconv"

	Game "github.com/kimulaco/trocon-server/domain/game"
)

type OwnedGame struct {
	AppId                    int    `json:"appId"`
	Name                     string `json:"name"`
	ImgIconUrl               string `json:"img_icon_url"`
	HasCommunityVisibleStats bool   `json:"has_community_visible_stats"`
	PlaytimeForever          int    `json:"playtime_forever"`
	PlaytimeWindowsForever   int    `json:"playtime_windows_forever"`
	PlaytimeMacForever       int    `json:"playtime_mac_forever"`
	PlaytimeLinuxForever     int    `json:"playtime_linux_forever"`
	RtimeLastPlayed          int    `json:"rtime_last_played"`
}

func (o OwnedGame) ToGame() Game.Game {
	appIdStr := strconv.Itoa(o.AppId)

	return Game.Game{
		AppId:                    o.AppId,
		Name:                     o.Name,
		IconImgUrl:               createIconImgUrl(appIdStr, o.ImgIconUrl),
		HeaderImgUrl:             createHeaderImgUrl(appIdStr),
		StoreUrl:                 createStoreUrl(appIdStr),
		HasCommunityVisibleStats: o.HasCommunityVisibleStats,
		Playtime:                 o.PlaytimeForever,
		RtimeLastPlayed:          o.RtimeLastPlayed,
	}
}

func createIconImgUrl(appId string, hash string) string {
	return "https://media.steampowered.com/steamcommunity/public/images/apps/" + appId + "/" + hash + ".jpg"
}

func createHeaderImgUrl(appId string) string {
	return "https://cdn.cloudflare.steamstatic.com/steam/apps/" + appId + "/header.jpg"
}

func createStoreUrl(appId string) string {
	return "https://store.steampowered.com/app/" + appId
}

func MapOwnedGamesToGames(ownedGames []OwnedGame) []Game.Game {
	games := make([]Game.Game, len(ownedGames))

	for i, ownedGame := range ownedGames {
		games[i] = ownedGame.ToGame()
	}

	return games
}

var TestGame1 = OwnedGame{
	AppId:                    1,
	Name:                     "Trophy Comp Game 1",
	ImgIconUrl:               "icon_hash",
	HasCommunityVisibleStats: true,
	PlaytimeForever:          1110,
	PlaytimeWindowsForever:   1000,
	PlaytimeMacForever:       100,
	PlaytimeLinuxForever:     10,
	RtimeLastPlayed:          0,
}

var TestGame2 = OwnedGame{
	AppId:                    2,
	Name:                     "Trophy Comp Game 2",
	ImgIconUrl:               "icon_hash2",
	HasCommunityVisibleStats: false,
	PlaytimeForever:          1000,
	PlaytimeWindowsForever:   1000,
	PlaytimeMacForever:       0,
	PlaytimeLinuxForever:     0,
	RtimeLastPlayed:          0,
}
