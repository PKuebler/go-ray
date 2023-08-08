package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONString(t *testing.T) {
	t.Parallel()

	testString := "{\"foo\":\"baa\"}"

	payload := NewJSONString(testString)
	assert.Implements(t, (*Payload)(nil), payload)

	assert.Equal(t, "json_string", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"value": testString,
	}, payload.GetContent())
}
