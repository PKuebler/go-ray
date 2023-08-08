package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplicationLog(t *testing.T) {
	t.Parallel()
	testString := "my custom log"

	payload := NewApplicationLog(testString)
	assert.Implements(t, (*Payload)(nil), payload)

	assert.Equal(t, "application_log", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"value": testString,
	}, payload.GetContent())
}
