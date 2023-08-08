package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {
	t.Parallel()

	testString := "red"

	var payload Payload
	payload = NewColor(testString)

	assert.Equal(t, "color", payload.GetType())
	assert.True(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"color": testString,
	}, payload.GetContent())
}
