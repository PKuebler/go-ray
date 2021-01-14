package payloads

import (
	"encoding/json"
)

type JSON struct {
	JSON interface{}
}

func NewJSON(json interface{}) JSON {
	return JSON{
		json,
	}
}

func (j JSON) GetType() string {
	return "json_string"
}

func (j JSON) GetContent() map[string]interface{} {
	jsonBytes, _ := json.Marshal(j.JSON)

	return map[string]interface{}{
		"value": string(jsonBytes),
	}
}

func (j JSON) ModifiesLastMessage() bool {
	return false
}
