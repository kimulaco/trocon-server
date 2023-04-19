package steamworks

type Player struct {
	SteamID                  string `json:"steamId"`
	CommunityVisibilityState int    `json:"communityVisibilityState"`
	ProfileState             int    `json:"profileState"`
	PersonaName              string `json:"personaName"`
	LastLogoff               int    `json:"lastLogoff"`
	ProfileUrl               string `json:"profileUrl"`
	Avatar                   string `json:"avatar"`
	AvatarMedium             string `json:"avatarMedium"`
	AvatarFull               string `json:"avatarFull"`
}

func (p *Player) IsEmpty() bool {
	return p.SteamID == "" ||
		p.PersonaName == "" ||
		p.ProfileUrl == ""
}

var TestUser = Player{
	SteamID:                  "0",
	CommunityVisibilityState: 3,
	ProfileState:             1,
	PersonaName:              "trophy-comp-user",
	LastLogoff:               1640962800,
	ProfileUrl:               "http://localhost:9999/id/0/",
	Avatar:                   "http://localhost:9999/avatar.jpg",
	AvatarMedium:             "http://localhost:9999/avatar_medium.jpg",
	AvatarFull:               "http://localhost:9999/avatar_full.jpg",
}
