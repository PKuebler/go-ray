package payloads

type Hide struct {
}

func NewHide() Hide {
	return Hide{}
}

func (h Hide) GetType() string {
	return "hide"
}

func (h Hide) GetContent() map[string]interface{} {
	return map[string]interface{}{}
}

func (h Hide) ModifiesLastMessage() bool {
	return true
}
