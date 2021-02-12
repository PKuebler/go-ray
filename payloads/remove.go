package payloads

// Remove payload
type Remove struct {
}

// NewRemove payload
func NewRemove() Remove {
	return Remove{}
}

// GetType of this payload
func (r Remove) GetType() string {
	return "remove"
}

// GetContent from this payload
func (r Remove) GetContent() map[string]interface{} {
	return map[string]interface{}{}
}

// ModifiesLastMessage by this payload?
func (r Remove) ModifiesLastMessage() bool {
	return true
}
