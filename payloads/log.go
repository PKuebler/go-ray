package payloads

type Log struct {
	Values []interface{}
}

func NewLog(values []interface{}) Log {
	return Log{
		values,
	}
}

func (l Log) GetType() string {
	return "log"
}

func (l Log) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"values": l.Values,
	}
}

func (l Log) ModifiesLastMessage() bool {
	return false
}
