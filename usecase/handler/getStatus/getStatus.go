package GetStatusAPI

import (
	"errors"
	"log"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/kimulaco/trocon-server/interface/steamworks"
	"github.com/kimulaco/trocon-server/pkg/httputil"
	"github.com/labstack/echo/v4"
)

type GetUserResponse struct {
	StatusCode int `json:"statusCode"`
}

func GetStatus(c echo.Context) error {
	s := steamworks.NewSteamworks()
	if s.InvalidEnv() {
		errorLog("STATUS_ENVIROMENT_ERROR: undefined steam API key")
		return c.JSON(httputil.NewError500("STATUS_ENVIROMENT_ERROR", ""))
	}

	return c.JSON(http.StatusOK, GetUserResponse{
		StatusCode: http.StatusOK,
	})
}

func errorLog(s string) {
	log.Print(s)
	sentry.CaptureException(errors.New(s))
}
