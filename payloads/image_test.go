package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImage(t *testing.T) {
	testPath := "https://localhost/image.png"

	var payload Payload
	payload = NewImage(testPath)

	assert.Equal(t, "custom", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"content": "<img src=\"" + testPath + "\" alt=\"\" />",
		"label":   "Image",
	}, payload.GetContent())
}
