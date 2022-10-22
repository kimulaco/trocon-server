package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetListenPort(t *testing.T) {
	t.Setenv("PORT", "")
	port1 := GetListenPort("")
	assert.Equal(t, ":", port1)

	t.Setenv("PORT", "9999")
	port2 := GetListenPort("3000")
	assert.Equal(t, ":9999", port2)

	t.Setenv("PORT", "")
	port3 := GetListenPort("3000")
	assert.Equal(t, ":3000", port3)
}
