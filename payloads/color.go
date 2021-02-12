package payloads

// Color payload
type Color struct {
	Color string
}

// NewColor payload
func NewColor(color string) Color {
	return Color{
		color,
	}
}

// GetType of this payload
func (c Color) GetType() string {
	return "color"
}

// GetContent from this payload
func (c Color) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"color": c.Color,
	}
}

// ModifiesLastMessage by this payload?
func (c Color) ModifiesLastMessage() bool {
	return true
}
