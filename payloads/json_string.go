package payloads

// JSONString payload
type JSONString struct {
	Value string
}

// NewJSONString payload
func NewJSONString(value string) JSONString {
	return JSONString{
		value,
	}
}

// GetType of this payload
func (j JSONString) GetType() string {
	return "json_string"
}

// GetContent from this payload
func (j JSONString) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"value": j.Value,
	}
}

// ModifiesLastMessage by this payload?
func (j JSONString) ModifiesLastMessage() bool {
	return false
}
