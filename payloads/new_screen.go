package payloads

type Screen struct {
	Name string
}

func NewScreen(name string) Screen {
	return Screen{
		name,
	}
}

func (n Screen) GetType() string {
	return "new_screen"
}

func (n Screen) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"name": n.Name,
	}
}

func (n Screen) ModifiesLastMessage() bool {
	return false
}
