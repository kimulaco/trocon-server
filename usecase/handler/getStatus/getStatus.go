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
		message := "STATUS_ENVIROMENT_ERROR: undefined steam API key"

		log.Print(message)
		sentry.CaptureException(errors.New(message))

		return c.JSON(httputil.NewError500("STATUS_ENVIROMENT_ERROR", ""))
	}

	return c.JSON(http.StatusOK, GetUserResponse{
		StatusCode: http.StatusOK,
	})
}
