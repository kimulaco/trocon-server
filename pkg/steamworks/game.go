package steamworks

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
