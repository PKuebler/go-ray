package payloads

// HideApp payload
type HideApp struct {
}

// NewHideApp payload
func NewHideApp() HideApp {
	return HideApp{}
}

// GetType of this payload
func (h HideApp) GetType() string {
	return "hide_app"
}

// GetContent from this payload
func (h HideApp) GetContent() map[string]interface{} {
	return map[string]interface{}{}
}

// ModifiesLastMessage by this payload?
func (h HideApp) ModifiesLastMessage() bool {
	return false
}
