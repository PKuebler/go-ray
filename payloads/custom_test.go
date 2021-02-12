package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustom(t *testing.T) {
	testContent := "my custom content"
	testLabel := "my test label"

	var payload Payload
	payload = NewCustom(testLabel, testContent)

	assert.Equal(t, "custom", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"content": testContent,
		"label":   testLabel,
	}, payload.GetContent())
}
