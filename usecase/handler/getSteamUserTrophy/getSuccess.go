package GetSteamUserTrophyAPI

import (
	"github.com/kimulaco/trocon-server/interface/steamworks"
	"github.com/kimulaco/trocon-server/pkg/stringsx"
)

func GetSuccess(game steamworks.GetPlayerAchievementsResponseOwnedGame) bool {
	if !game.Success && stringsx.Contains(game.Error, NON_ERROR_MESSAGES) {
		return true
	}
	return game.Success
}

var NON_ERROR_MESSAGES = []string{
	"Requested app has no stats",
}
