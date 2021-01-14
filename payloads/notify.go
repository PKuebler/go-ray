package payloads

type Notify struct {
	Text string
}

func NewNotify(text string) Notify {
	return Notify{
		text,
	}
}

func (n Notify) GetType() string {
	return "notify"
}

func (n Notify) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"value": n.Text,
	}
}

func (n Notify) ModifiesLastMessage() bool {
	return false
}
