package ray

// Message to communicate with the ray app
type Message struct {
	UUID     string                   `json:"uuid"`
	Payloads []map[string]interface{} `json:"payloads"`
	Meta     map[string]string        `json:"meta"`
}
