package payloads

type ApplicationLog struct {
	ApplicationLog string
}

func NewApplicationLog(log string) ApplicationLog {
	return ApplicationLog{
		log,
	}
}

func (a ApplicationLog) GetType() string {
	return "application_log"
}

func (a ApplicationLog) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"value": a.ApplicationLog,
	}
}

func (a ApplicationLog) ModifiesLastMessage() bool {
	return false
}
