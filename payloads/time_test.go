package payloads

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	testTime := time.Now()

	var payload Payload
	payload = NewTime(testTime)

	assert.Equal(t, "carbon", payload.GetType())
	assert.False(t, payload.ModifiesLastMessage())
	assert.Equal(t, map[string]interface{}{
		"formatted": testTime.Format(time.RFC1123),
		"timestamp": testTime.Unix(),
		"timezone":  testTime.Location(),
	}, payload.GetContent())
}
