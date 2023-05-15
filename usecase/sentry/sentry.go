package sentry

import (
	"errors"
	"os"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
)

func Init() error {
	dsn := os.Getenv("SENTRY_DSN")

	if dsn == "" {
		return errors.New("SENTRY_DSN undefined")
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		TracesSampleRate: 1.0,
	})

	if err != nil {
		return err
	}

	return nil
}

func Middleware() echo.MiddlewareFunc {
	return sentryecho.New(sentryecho.Options{
		Repanic: true,
	})
}

func SetTag(c echo.Context, tag string, v string) {
	if hub := sentryecho.GetHubFromContext(c); hub != nil {
		hub.Scope().SetTag(tag, v)
	}
}
