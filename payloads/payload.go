package payloads

// Payload of a ray message
type Payload interface {
	// GetType of this payload
	GetType() string
	// GetContent from this payload
	GetContent() map[string]interface{}
	// ModifiesLastMessage by this payload?
	ModifiesLastMessage() bool
}
