package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotify(t *testing.T) {
	testString := "my custom notification"

	var payload Payload
	payload = NewNotify(testString)

	assert.Equal(t, "notify", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"value": testString,
	}, payload.GetContent())
}
