package testdata

import (
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func InitEcho(path string, query string) (*httptest.ResponseRecorder, echo.Context) {
	target := path
	if query != "" {
		target += "?" + query
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, target, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(path)

	return rec, c
}
