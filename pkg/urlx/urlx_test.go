package urlx

import (
	"errors"
	"strconv"
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

func BenchmarkUrlx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pageStr := strconv.Itoa(i)
		nextPageStr := strconv.Itoa(i + 1)
		prevPageStr := strconv.Itoa(i - 1)
		expect := "http://localhost?page=" + pageStr + "&nextPage=" + nextPageStr + "&prevPage=" + prevPageStr

		u, err := NewUrlx("http://localhost")

		if err != nil {
			b.Fatalf("Error during request: %s", err)
		}

		u.AddQuery("page", pageStr)
		u.AddQuery("nextPage", nextPageStr)
		u.AddQuery("prevPage", prevPageStr)

		if u.ToString() != expect {
			b.Fatalf("Unexpected status u.ToString(): %s", u.ToString())
		}
	}
}