package payloads

type Remove struct {
}

func NewRemove() Remove {
	return Remove{}
}

func (r Remove) GetType() string {
	return "remove"
}

func (r Remove) GetContent() map[string]interface{} {
	return map[string]interface{}{}
}

func (r Remove) ModifiesLastMessage() bool {
	return true
}
