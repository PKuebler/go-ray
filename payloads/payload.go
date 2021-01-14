package payloads

// Payload of a ray message
type Payload interface {
	GetType() string
	GetContent() map[string]interface{}
	ModifiesLastMessage() bool
}
