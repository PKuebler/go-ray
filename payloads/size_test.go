package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	t.Parallel()

	testFontSize := "lg"

	payload := NewSize(testFontSize)
	assert.Implements(t, (*Payload)(nil), payload)

	assert.Equal(t, "size", payload.GetType())
	assert.True(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"size": testFontSize,
	}, payload.GetContent())
}
