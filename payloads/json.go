package payloads

import (
	"encoding/json"
)

// JSON payload
type JSON struct {
	JSON interface{}
}

// NewJSON payload
func NewJSON(json interface{}) JSON {
	return JSON{
		json,
	}
}

// GetType of this payload
func (j JSON) GetType() string {
	return "json_string"
}

// GetContent from this payload
func (j JSON) GetContent() map[string]interface{} {
	jsonBytes, _ := json.Marshal(j.JSON)

	return map[string]interface{}{
		"value": string(jsonBytes),
	}
}

// ModifiesLastMessage by this payload?
func (j JSON) ModifiesLastMessage() bool {
	return false
}
