package testdata

import (
	"github.com/kimulaco/trophy-comp-server/pkg/steamworks"
)

var TestAchievement1 = steamworks.Achievement{
	ApiName:     "api-1",
	Name:        "Trophy 1",
	Description: "",
	Achieved:    1,
	UnlockTime:  1640962800,
}

var TestAchievement2 = steamworks.Achievement{
	ApiName:     "api-2",
	Name:        "Trophy 2",
	Description: "",
	Achieved:    0,
	UnlockTime:  1640962800,
}
