package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	t.Parallel()

	var payload Payload
	payload = NewBool(true)

	assert.Equal(t, "custom", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"content": true,
		"label":   "Boolean",
	}, payload.GetContent())
}
