package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHideApp(t *testing.T) {
	t.Parallel()

	payload := NewHideApp()
	assert.Implements(t, (*Payload)(nil), payload)

	assert.Equal(t, "hide_app", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{}, payload.GetContent())
}
