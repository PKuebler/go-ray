package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTML(t *testing.T) {
	t.Parallel()

	testString := "<a href=\"#\">MyLink</a>"

	payload := NewHTML(testString)
	assert.Implements(t, (*Payload)(nil), payload)

	assert.Equal(t, "custom", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"content": testString,
		"label":   "HTML",
	}, payload.GetContent())
}
