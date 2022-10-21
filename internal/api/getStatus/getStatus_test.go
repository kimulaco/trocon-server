package GetStatusAPI

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	status200 string = `{"statusCode":200}`
	envError  string = `{"statusCode":500,"errorCode":"STATUS_ENVIROMENT_ERROR","message":"internal server error"}`
)

func TestGetStatus200(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetStatus(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, status200, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestGetStatusEnvError(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "")

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetStatus(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, envError, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}
