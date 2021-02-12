package payloads

// ClearAll payload
type ClearAll struct {
}

// NewClearAll payload
func NewClearAll() ClearAll {
	return ClearAll{}
}

// GetType of this payload
func (c ClearAll) GetType() string {
	return "clear_all"
}

// GetContent from this payload
func (c ClearAll) GetContent() map[string]interface{} {
	return map[string]interface{}{}
}

// ModifiesLastMessage by this payload?
func (c ClearAll) ModifiesLastMessage() bool {
	return false
}
