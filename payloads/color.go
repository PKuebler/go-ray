package payloads

type Color struct {
	Color string
}

func NewColor(color string) Color {
	return Color{
		color,
	}
}

func (c Color) GetType() string {
	return "color"
}

func (c Color) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"color": c.Color,
	}
}

func (c Color) ModifiesLastMessage() bool {
	return true
}
