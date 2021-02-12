package payloads

// Bool payload
type Bool struct {
	Bool bool
}

// NewBool payload
func NewBool(b bool) Bool {
	return Bool{
		Bool: b,
	}
}

// GetType of this payload
func (b Bool) GetType() string {
	return "custom"
}

// GetContent from this payload
func (b Bool) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"content": b.Bool,
		"label":   "Boolean",
	}
}

// ModifiesLastMessage by this payload?
func (b Bool) ModifiesLastMessage() bool {
	return false
}
