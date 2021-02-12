package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTML(t *testing.T) {
	testString := "<a href=\"#\">MyLink</a>"

	var payload Payload
	payload = NewHTML(testString)

	assert.Equal(t, "custom", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"content": testString,
		"label":   "HTML",
	}, payload.GetContent())
}
