package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHideApp(t *testing.T) {
	var payload Payload
	payload = NewHideApp()

	assert.Equal(t, "hide_app", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{}, payload.GetContent())
}
