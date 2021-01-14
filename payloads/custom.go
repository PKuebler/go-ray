package payloads

type Custom struct {
	Content string
	Label   string
}

func NewCustom(label string, content string) Custom {
	return Custom{
		content,
		label,
	}
}

func (c Custom) GetType() string {
	return "custom"
}

func (c Custom) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"content": c.Content,
		"label":   c.Label,
	}
}

func (c Custom) ModifiesLastMessage() bool {
	return false
}
