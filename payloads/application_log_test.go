package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplicationLog(t *testing.T) {
	testString := "my custom log"

	var payload Payload
	payload = NewApplicationLog(testString)

	assert.Equal(t, "application_log", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"value": testString,
	}, payload.GetContent())
}
