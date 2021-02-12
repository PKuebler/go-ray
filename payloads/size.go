package payloads

// Size payload
type Size struct {
	Size string
}

// NewSize payload
func NewSize(size string) Size {
	return Size{
		size,
	}
}

// GetType of this payload
func (s Size) GetType() string {
	return "size"
}

// GetContent from this payload
func (s Size) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"size": s.Size,
	}
}

// ModifiesLastMessage by this payload?
func (s Size) ModifiesLastMessage() bool {
	return true
}
