package payloads

import "time"

// Time payload
type Time struct {
	Time time.Time
}

// NewTime payload
func NewTime(t time.Time) Time {
	return Time{
		Time: t,
	}
}

// GetType of this payload
func (t Time) GetType() string {
	return "carbon"
}

// GetContent from this payload
func (t Time) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"formatted": t.Time.Format(time.RFC1123),
		"timestamp": t.Time.Unix(),
		"timezone":  t.Time.Location(),
	}
}

// ModifiesLastMessage by this payload?
func (t Time) ModifiesLastMessage() bool {
	return false
}
