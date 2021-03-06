package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	testFontSize := "lg"

	var payload Payload
	payload = NewSize(testFontSize)

	assert.Equal(t, "size", payload.GetType())
	assert.True(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"size": testFontSize,
	}, payload.GetContent())
}
