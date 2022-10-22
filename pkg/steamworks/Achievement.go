package steamworks

type Achievement struct {
	ApiName     string `json:"apiname"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Achieved    int    `json:"achieved"`
	UnlockTime  int    `json:"unlocktime"`
}
