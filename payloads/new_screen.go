package payloads

// Screen payload
type Screen struct {
	Name string
}

// NewScreen payload
func NewScreen(name string) Screen {
	return Screen{
		name,
	}
}

// GetType of this payload
func (n Screen) GetType() string {
	return "new_screen"
}

// GetContent from this payload
func (n Screen) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"name": n.Name,
	}
}

// ModifiesLastMessage by this payload?
func (n Screen) ModifiesLastMessage() bool {
	return false
}
