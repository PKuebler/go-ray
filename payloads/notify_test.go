package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotify(t *testing.T) {
	t.Parallel()

	testString := "my custom notification"

	payload := NewNotify(testString)
	assert.Implements(t, (*Payload)(nil), payload)

	assert.Equal(t, "notify", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"value": testString,
	}, payload.GetContent())
}
