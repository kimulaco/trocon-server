package steamworks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSteamworks(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "XXXXXXXX")
	t.Setenv("STEAM_API_BASE_URL", "http://localhost:9999")

	s := NewSteamworks()

	assert.Equal(t, s.APIKey, "XXXXXXXX")
	assert.Equal(t, s.APIBaseURL, "http://localhost:9999")
	assert.Equal(t, s.InvalidEnv(), false)
}

func TestInvalidEnv(t *testing.T) {
	t.Setenv("STEAM_API_KEY", "")
	t.Setenv("STEAM_API_BASE_URL", "")

	s := NewSteamworks()

	assert.Equal(t, s.APIKey, "")
	assert.Equal(t, s.APIBaseURL, "")
	assert.Equal(t, s.InvalidEnv(), true)
}

var Response403 = `<html>

<head>
	<title>Forbidden</title>
</head>

<body>
	<h1>Forbidden</h1>Access is denied. Retrying will not help. Please verify your
	<pre>key=</pre> parameter.
</body>

</html>`
