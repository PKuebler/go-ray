package payloads

type CreateLock struct {
	Name string
}

func NewCreateLock(name string) CreateLock {
	return CreateLock{
		name,
	}
}

func (c CreateLock) GetType() string {
	return "create_lock"
}

func (c CreateLock) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"name": c.Name,
	}
}

func (c CreateLock) ModifiesLastMessage() bool {
	return false
}
