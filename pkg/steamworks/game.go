package steamworks

import (
	"strconv"
)

type Game struct {
	AppId                    int    `json:"appId"`
	Name                     string `json:"name"`
	IconImgUrl               string `json:"iconImgUrl"`
	HeaderImgUrl             string `json:"headerImgUrl"`
	StoreUrl                 string `json:"storeUrl"`
	HasCommunityVisibleStats bool   `json:"hasCommunityVisibleStats"`
	Playtime                 int    `json:"playtime"`
	RtimeLastPlayed          int    `json:"rtimeLastPlayed"`
}

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

func (o OwnedGame) ToGame() Game {
	appIdStr := strconv.Itoa(o.AppId)

	return Game{
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

func MapOwnedGamesToGames(ownedGames []OwnedGame) []Game {
	games := make([]Game, len(ownedGames))

	for i, ownedGame := range ownedGames {
		games[i] = ownedGame.ToGame()
	}

	return games
}
