package payloads

type Size struct {
	Size string
}

func NewSize(size string) Size {
	return Size{
		size,
	}
}

func (s Size) GetType() string {
	return "size"
}

func (s Size) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"size": s.Size,
	}
}

func (s Size) ModifiesLastMessage() bool {
	return true
}
