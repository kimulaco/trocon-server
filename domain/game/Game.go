package Game

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
