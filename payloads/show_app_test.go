package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowApp(t *testing.T) {
	t.Parallel()

	var payload Payload
	payload = NewShowApp()

	assert.Equal(t, "show_app", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{}, payload.GetContent())
}
