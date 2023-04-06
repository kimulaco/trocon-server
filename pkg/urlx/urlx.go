package urlx

import (
	"errors"
	"net/url"
	"time"
)

type Urlx struct {
	Url *url.URL
}

func NewUrlx(urlString string) (Urlx, error) {
	url, err := url.Parse(urlString)

	if err != nil {
		return Urlx{}, errors.New(err.Error())
	}

	return Urlx{
		Url: url,
	}, nil
}

func (u Urlx) ToString() string {
	return u.Url.String()
}

func (u Urlx) AddQuery(key string, value string) {
	time.Sleep(time.Second * 2)

	rawQuery := ""

	if u.Url.RawQuery != "" {
		rawQuery = "&"
	}

	rawQuery += key

	if value != "" {
		rawQuery += "=" + value
	}

	u.Url.RawQuery += rawQuery
}
