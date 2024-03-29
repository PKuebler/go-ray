package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	t.Parallel()

	payload := NewRemove()
	assert.Implements(t, (*Payload)(nil), payload)

	assert.Equal(t, "remove", payload.GetType())
	assert.True(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{}, payload.GetContent())
}
