package urlx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlx(t *testing.T) {
	u, err := NewUrlx("http://localhost?id=1")
	assert.Equal(t, "http://localhost?id=1", u.Url.String())
	assert.Equal(t, "http://localhost?id=1", u.ToString())
	assert.Equal(t, nil, err)

	u.AddQuery("page", "2")
	u.AddQuery("orderBy", "asc")
	assert.Equal(t, "http://localhost?id=1&page=2&orderBy=asc", u.Url.String())
	assert.Equal(t, "http://localhost?id=1&page=2&orderBy=asc", u.ToString())
}

func TestUrlxError(t *testing.T) {
	u, err := NewUrlx("http://\\\\localhost")
	assert.Equal(t, Urlx{}, u)
	assert.Equal(t, errors.New("parse \"http://\\\\\\\\localhost\": invalid character \"\\\\\" in host name"), err)
}
