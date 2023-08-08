package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClearAll(t *testing.T) {
	t.Parallel()

	payload := NewClearAll()
	assert.Implements(t, (*Payload)(nil), payload)

	assert.Equal(t, "clear_all", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{}, payload.GetContent())
}
