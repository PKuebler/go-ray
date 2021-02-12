package payloads

// Custom payload
type Custom struct {
	Content string
	Label   string
}

// NewCustom payload
func NewCustom(label string, content string) Custom {
	return Custom{
		content,
		label,
	}
}

// GetType of this payload
func (c Custom) GetType() string {
	return "custom"
}

// GetContent from this payload
func (c Custom) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"content": c.Content,
		"label":   c.Label,
	}
}

// ModifiesLastMessage by this payload?
func (c Custom) ModifiesLastMessage() bool {
	return false
}
