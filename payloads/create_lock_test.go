package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLock(t *testing.T) {
	testString := "myName"

	var payload Payload
	payload = NewCreateLock(testString)

	assert.Equal(t, "create_lock", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"name": testString,
	}, payload.GetContent())
}
