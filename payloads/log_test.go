package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	t.Parallel()

	testValues := []interface{}{
		"my custom log",
		false,
		3132,
		uint32(23234),
	}

	payload := NewLog(testValues)
	assert.Implements(t, (*Payload)(nil), payload)

	assert.Equal(t, "log", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"values": testValues,
	}, payload.GetContent())
}
