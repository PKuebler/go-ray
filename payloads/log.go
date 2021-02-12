package payloads

// Log payload
type Log struct {
	Values []interface{}
}

// NewLog payload
func NewLog(values []interface{}) Log {
	return Log{
		values,
	}
}

// GetType of this payload
func (l Log) GetType() string {
	return "log"
}

// GetContent from this payload
func (l Log) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"values": l.Values,
	}
}

// ModifiesLastMessage by this payload?
func (l Log) ModifiesLastMessage() bool {
	return false
}
