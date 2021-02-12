package payloads

// CreateLock payload
type CreateLock struct {
	Name string
}

// NewCreateLock payload
func NewCreateLock(name string) CreateLock {
	return CreateLock{
		name,
	}
}

// GetType of this payload
func (c CreateLock) GetType() string {
	return "create_lock"
}

// GetContent from this payload
func (c CreateLock) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"name": c.Name,
	}
}

// ModifiesLastMessage by this payload?
func (c CreateLock) ModifiesLastMessage() bool {
	return false
}
