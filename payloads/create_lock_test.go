package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLock(t *testing.T) {
	t.Parallel()

	testString := "myName"

	payload := NewCreateLock(testString)
	assert.Implements(t, (*Payload)(nil), payload)

	assert.Equal(t, "create_lock", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"name": testString,
	}, payload.GetContent())
}
