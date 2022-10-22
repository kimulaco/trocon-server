package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCORSAllowOriginsDev(t *testing.T) {
	t.Setenv("BUILD_MODE", "development")
	t.Setenv("CORS_ALLOW_ORIGIN", "http://example.com")
	origins1 := GetCORSAllowOrigins()
	assert.Equal(t, []string{"http://localhost:*"}, origins1)
}

func TestGetCORSAllowOriginsProd(t *testing.T) {
	t.Setenv("BUILD_MODE", "production")
	t.Setenv("CORS_ALLOW_ORIGIN", "http://example.com")
	origins1 := GetCORSAllowOrigins()
	assert.Equal(t, []string{"http://example.com"}, origins1)

	t.Setenv("CORS_ALLOW_ORIGIN", "")
	origins2 := GetCORSAllowOrigins()
	assert.Equal(t, []string{}, origins2)
}

func TestGetCORSAllowOriginsNoBuildMode(t *testing.T) {
	t.Setenv("BUILD_MODE", "")
	t.Setenv("CORS_ALLOW_ORIGIN", "http://example.com")
	origins1 := GetCORSAllowOrigins()
	assert.Equal(t, []string{}, origins1)
}
