package payloads

// ApplicationLog payload
type ApplicationLog struct {
	ApplicationLog string
}

// NewApplicationLog payload
func NewApplicationLog(log string) ApplicationLog {
	return ApplicationLog{
		log,
	}
}

// GetType of this payload
func (a ApplicationLog) GetType() string {
	return "application_log"
}

// GetContent from this payload
func (a ApplicationLog) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"value": a.ApplicationLog,
	}
}

// ModifiesLastMessage by this payload?
func (a ApplicationLog) ModifiesLastMessage() bool {
	return false
}
