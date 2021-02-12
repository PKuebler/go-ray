package payloads

// Hide payload
type Hide struct {
}

// NewHide payload
func NewHide() Hide {
	return Hide{}
}

// GetType of this payload
func (h Hide) GetType() string {
	return "hide"
}

// GetContent from this payload
func (h Hide) GetContent() map[string]interface{} {
	return map[string]interface{}{}
}

// ModifiesLastMessage by this payload?
func (h Hide) ModifiesLastMessage() bool {
	return true
}
