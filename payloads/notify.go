package payloads

// Notify payload
type Notify struct {
	Text string
}

// NewNotify payload
func NewNotify(text string) Notify {
	return Notify{
		text,
	}
}

// GetType of this payload
func (n Notify) GetType() string {
	return "notify"
}

// GetContent from this payload
func (n Notify) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"value": n.Text,
	}
}

// ModifiesLastMessage by this payload?
func (n Notify) ModifiesLastMessage() bool {
	return false
}
