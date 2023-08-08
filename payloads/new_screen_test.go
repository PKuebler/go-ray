package payloads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewScreen(t *testing.T) {
	t.Parallel()

	testString := "my screen name"

	var payload Payload
	payload = NewScreen(testString)

	assert.Equal(t, "new_screen", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"name": testString,
	}, payload.GetContent())
}
