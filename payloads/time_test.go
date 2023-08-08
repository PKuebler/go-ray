package payloads

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	t.Parallel()

	testTime := time.Now()

	payload := NewTime(testTime)
	assert.Implements(t, (*Payload)(nil), payload)

	assert.Equal(t, "carbon", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"formatted": testTime.Format(time.RFC1123),
		"timestamp": testTime.Unix(),
		"timezone":  testTime.Location(),
	}, payload.GetContent())
}
