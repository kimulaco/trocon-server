package steamworks

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/kimulaco/trocon-server/pkg/httputil"
	"github.com/kimulaco/trocon-server/pkg/urlx"
)

type ResolveVanityURLResponse struct {
	Response struct {
		Steamid string `json:"steamid"`
		Success int    `json:"success"`
		Message string `json:"message"`
	} `json:"response"`
}

const ResolveVanityURLPath = "/ISteamUser/ResolveVanityURL/v1/"

func (s Steamworks) ResolveVanityURL(vanityurl string) (string, error) {
	if vanityurl == "" {
		return "", errors.New("vanityURL is required")
	}

	apiUrl, _ := urlx.NewUrlx(s.APIBaseURL + ResolveVanityURLPath)
	apiUrl.AddQuery("key", s.APIKey)
	apiUrl.AddQuery("vanityurl", vanityurl)

	res, err := http.Get(apiUrl.ToString())
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", errors.New(res.Status)
	}

	resBody, err := httputil.ReadBody[ResolveVanityURLResponse](res)
	if err != nil {
		return "", err
	}
	if resBody.Response.Success != 1 {
		return "", errors.New("success:" + strconv.Itoa(resBody.Response.Success) + " " + resBody.Response.Message)
	}

	return resBody.Response.Steamid, nil
}
