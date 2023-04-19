package steamworks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer(t *testing.T) {
	p := Player{
		SteamID:     "1",
		PersonaName: "name",
		ProfileUrl:  "http://localhost:9999",
	}

	assert.Equal(t, p.IsEmpty(), false)
}

func TestPlayerEmpty(t *testing.T) {
	p := Player{}

	assert.Equal(t, p.IsEmpty(), true)
}
