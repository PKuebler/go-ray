package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	t.Parallel()

	payload := NewBool(true)
	assert.Implements(t, (*Payload)(nil), payload)

	assert.Equal(t, "custom", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"content": true,
		"label":   "Boolean",
	}, payload.GetContent())
}
