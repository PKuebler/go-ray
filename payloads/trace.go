package payloads

// Trace payload
type Trace struct {
	FileName   string
	LineNumber int
	ClassName  string
	MethodName string
}

// NewTrace payload
func NewTrace() Trace {
	return Trace{
		"filename.go",
		203,
		"className",
		"methodName",
	}
}

// GetType of this payload
func (t Trace) GetType() string {
	return "trace"
}

// GetContent from this payload
func (t Trace) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"frames": []map[string]interface{}{
			{
				"file_name":   t.FileName,
				"line_number": t.LineNumber,
				"class":       t.ClassName,
				"method":      t.MethodName,
			},
		},
	}
}

// ModifiesLastMessage by this payload?
func (t Trace) ModifiesLastMessage() bool {
	return false
}
