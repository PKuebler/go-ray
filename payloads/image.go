package payloads

import "fmt"

// Image payload
type Image struct {
	Path string
}

// NewImage payload
func NewImage(path string) Image {
	return Image{
		Path: path,
	}
}

// GetType of this payload
func (i Image) GetType() string {
	return "custom"
}

// GetContent from this payload
func (i Image) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"content": fmt.Sprintf("<img src=\"%s\" alt=\"\" />", i.Path),
		"label":   "Image",
	}
}

// ModifiesLastMessage by this payload?
func (i Image) ModifiesLastMessage() bool {
	return false
}
