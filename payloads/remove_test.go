package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	t.Parallel()

	var payload Payload
	payload = NewRemove()

	assert.Equal(t, "remove", payload.GetType())
	assert.True(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{}, payload.GetContent())
}
