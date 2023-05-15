package sentry

import (
	"errors"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
)

func Init() error {
	dsn := os.Getenv("SENTRY_DSN")

	if dsn == "" {
		return errors.New("SENTRY_DSN undefined")
	}

	env := os.Getenv("BUILD_MODE")

	if env == "" {
		return errors.New("BUILD_MODE undefined")
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		TracesSampleRate: 1.0,
		Environment:      env,
	})

	if err != nil {
		return err
	}

	return nil
}

func NewMiddleware() echo.MiddlewareFunc {
	return sentryecho.New(sentryecho.Options{
		Repanic: true,
	})
}

func SetTag(c echo.Context, tag string, v string) {
	if hub := sentryecho.GetHubFromContext(c); hub != nil {
		hub.Scope().SetTag(tag, v)
	}
}

func AddBreadcrumb(s string) {
	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Message: s,
	})
}

func Recover() {
	sentry.Recover()
	sentry.Flush(2 * time.Second)
}
