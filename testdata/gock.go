package testdata

import (
	"github.com/h2non/gock"
)

type GockConfig struct {
	Url      string
	Path     string
	Querys   map[string]string
	Response interface{}
}

func InitGock(config GockConfig) (*gock.Request, *gock.Response) {
	req := gock.New(config.Url).Get(config.Path)

	if len(config.Querys) > 0 {
		for key, value := range config.Querys {
			req.MatchParam(key, value)
		}
	}

	res := req.Reply(200).JSON(config.Response)

	return req, res
}
