package urlx

import (
	"net/url"
)

type Urlx struct {
	Url *url.URL
}

func NewUrlx(urlString string) (Urlx, error) {
	url, err := url.Parse(urlString)

	if err != nil {
		return Urlx{}, err
	}

	return Urlx{
		Url: url,
	}, nil
}

func (u Urlx) ToString() string {
	return u.Url.String()
}

func (u Urlx) AddQuery(key string, value string) {
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
