package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustom(t *testing.T) {
	t.Parallel()

	testContent := "my custom content"
	testLabel := "my test label"

	payload := NewCustom(testLabel, testContent)
	assert.Implements(t, (*Payload)(nil), payload)

	assert.Equal(t, "custom", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"content": testContent,
		"label":   testLabel,
	}, payload.GetContent())
}
