package steamworks

import (
	"errors"
	"testing"

	"github.com/h2non/gock"
	"github.com/kimulaco/trocon-server/testdata"
	"github.com/stretchr/testify/assert"
)

func TestResolveVanityURL(t *testing.T) {
	const baseUrl = "http://localhost:9999"
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", baseUrl)
	s := NewSteamworks()

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:        baseUrl,
		Path:       ResolveVanityURLPath,
		StatusCode: 200,
		Response:   ResolveVanityURLResponse200,
	})

	steamid, err := s.ResolveVanityURL("trophy-comp-user")
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, steamid, "1234567890123")
}

func TestResolveVanityURLError_NoMatch(t *testing.T) {
	const baseUrl = "http://localhost:9999"
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", baseUrl)
	s := NewSteamworks()

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:        baseUrl,
		Path:       ResolveVanityURLPath,
		StatusCode: 200,
		Response:   ResolveVanityURLResponse200_43,
	})

	steamid1, err1 := s.ResolveVanityURL("undefined-user-name")
	assert.Equal(t, steamid1, "")
	assert.Equal(t, err1, errors.New("success:42 No match"))
}

func TestResolveVanityURLError_Forbidden(t *testing.T) {
	const baseUrl = "http://localhost:9999"
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", baseUrl)
	s := NewSteamworks()

	defer gock.Off()
	testdata.InitGock(testdata.GockConfig{
		Url:        baseUrl,
		Path:       ResolveVanityURLPath,
		StatusCode: 403,
		Response:   Response403,
	})

	steamid2, err2 := s.ResolveVanityURL("forbidden-user-name")
	assert.Equal(t, steamid2, "")
	assert.Equal(t, err2, errors.New("403 Forbidden"))
}

var ResolveVanityURLResponse200 = ResolveVanityURLResponse{
	Response: struct {
		Steamid string `json:"steamid"`
		Success int    `json:"success"`
		Message string `json:"message"`
	}{
		Steamid: "1234567890123",
		Success: 1,
		Message: "",
	},
}

var ResolveVanityURLResponse200_43 = ResolveVanityURLResponse{
	Response: struct {
		Steamid string `json:"steamid"`
		Success int    `json:"success"`
		Message string `json:"message"`
	}{
		Steamid: "",
		Success: 42,
		Message: "No match",
	},
}
