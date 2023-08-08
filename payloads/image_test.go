package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImage(t *testing.T) {
	t.Parallel()

	testPath := "https://localhost/image.png"

	payload := NewImage(testPath)
	assert.Implements(t, (*Payload)(nil), payload)

	assert.Equal(t, "custom", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"content": "<img src=\"" + testPath + "\" alt=\"\" />",
		"label":   "Image",
	}, payload.GetContent())
}
