package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONString(t *testing.T) {
	testString := "{\"foo\":\"baa\"}"

	var payload Payload
	payload = NewJSONString(testString)

	assert.Equal(t, "json_string", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"value": testString,
	}, payload.GetContent())
}
