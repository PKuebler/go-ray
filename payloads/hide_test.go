package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHide(t *testing.T) {
	t.Parallel()

	var payload Payload
	payload = NewHide()

	assert.Equal(t, "hide", payload.GetType())
	assert.True(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{}, payload.GetContent())
}
