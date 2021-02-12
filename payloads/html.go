package payloads

// HTML payload
type HTML struct {
	HTML string
}

// NewHTML payload
func NewHTML(h string) HTML {
	return HTML{
		HTML: h,
	}
}

// GetType of this payload
func (h HTML) GetType() string {
	return "custom"
}

// GetContent from this payload
func (h HTML) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"content": h.HTML,
		"label":   "HTML",
	}
}

// ModifiesLastMessage by this payload?
func (h HTML) ModifiesLastMessage() bool {
	return false
}
