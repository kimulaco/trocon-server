package GetSteamUserTrophyAPI

import (
	"testing"

	"github.com/kimulaco/trocon-server/interface/steamworks"
	"github.com/stretchr/testify/assert"
)

func TestGetSuccess(t *testing.T) {
	assert.Equal(t, GetSuccess(SuccessGame), true)
	assert.Equal(t, GetSuccess(NoStatsGame), true)
	assert.Equal(t, GetSuccess(NonSuccessGame), false)
}

var SuccessGame = steamworks.GetPlayerAchievementsResponseOwnedGame{
	Success: true,
	Error: "",
}

var NoStatsGame = steamworks.GetPlayerAchievementsResponseOwnedGame{
	Success: false,
	Error: "Requested app has no stats",
}

var NonSuccessGame = steamworks.GetPlayerAchievementsResponseOwnedGame{
	Success: false,
	Error: "",
}
