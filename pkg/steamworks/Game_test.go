package steamworks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOwnedGameToGame(t *testing.T) {
	g := TestGame1

	assert.Equal(t, g.ToGame(), Game{
		AppId:                    1,
		Name:                     "Trophy Comp Game 1",
		IconImgUrl:               "https://media.steampowered.com/steamcommunity/public/images/apps/1/icon_hash.jpg",
		HeaderImgUrl:             "https://cdn.cloudflare.steamstatic.com/steam/apps/1/header.jpg",
		StoreUrl:                 "https://store.steampowered.com/app/1",
		HasCommunityVisibleStats: true,
		Playtime:                 1110,
		RtimeLastPlayed:          0,
	})
}

func TestMapOwnedGamesToGames(t *testing.T) {
	games := MapOwnedGamesToGames([]OwnedGame{TestGame1, TestGame2})

	assert.Equal(t, games, []Game{
		{
			AppId:                    1,
			Name:                     "Trophy Comp Game 1",
			IconImgUrl:               "https://media.steampowered.com/steamcommunity/public/images/apps/1/icon_hash.jpg",
			HeaderImgUrl:             "https://cdn.cloudflare.steamstatic.com/steam/apps/1/header.jpg",
			StoreUrl:                 "https://store.steampowered.com/app/1",
			HasCommunityVisibleStats: true,
			Playtime:                 1110,
			RtimeLastPlayed:          0,
		},
		{
			AppId:                    2,
			Name:                     "Trophy Comp Game 2",
			IconImgUrl:               "https://media.steampowered.com/steamcommunity/public/images/apps/2/icon_hash2.jpg",
			HeaderImgUrl:             "https://cdn.cloudflare.steamstatic.com/steam/apps/2/header.jpg",
			StoreUrl:                 "https://store.steampowered.com/app/2",
			HasCommunityVisibleStats: false,
			Playtime:                 1000,
			RtimeLastPlayed:          0,
		},
	})
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
