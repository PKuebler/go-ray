package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowApp(t *testing.T) {
	t.Parallel()

	payload := NewShowApp()
	assert.Implements(t, (*Payload)(nil), payload)

	assert.Equal(t, "show_app", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{}, payload.GetContent())
}
