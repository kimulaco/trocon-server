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
