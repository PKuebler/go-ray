package payloads

type JSONString struct {
	Value string
}

func NewJSONString(value string) JSONString {
	return JSONString{
		value,
	}
}

func (j JSONString) GetType() string {
	return "json_string"
}

func (j JSONString) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"value": j.Value,
	}
}

func (j JSONString) ModifiesLastMessage() bool {
	return false
}
