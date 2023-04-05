package testdata

import (
	"github.com/h2non/gock"
)

type GockConfig struct {
	Url      string
	Path     string
	Querys   map[string]string
	StatusCode int
	Response interface{}
}

func InitGock(config GockConfig) (*gock.Request, *gock.Response) {
	req := gock.New(config.Url).Get(config.Path).Persist()

	if len(config.Querys) > 0 {
		for key, value := range config.Querys {
			req.MatchParam(key, value)
		}
	}

	if config.StatusCode <= 0 {
		config.StatusCode = 200
	}

	res := req.Reply(config.StatusCode).JSON(config.Response)

	return req, res
}
