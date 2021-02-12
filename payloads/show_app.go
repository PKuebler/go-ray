package payloads

// ShowApp payload
type ShowApp struct {
}

// NewShowApp payload
func NewShowApp() ShowApp {
	return ShowApp{}
}

// GetType of this payload
func (s ShowApp) GetType() string {
	return "show_app"
}

// GetContent from this payload
func (s ShowApp) GetContent() map[string]interface{} {
	return map[string]interface{}{}
}

// ModifiesLastMessage by this payload?
func (s ShowApp) ModifiesLastMessage() bool {
	return false
}
