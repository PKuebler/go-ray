package payloads

// Caller payload
type Caller struct {
	FileName   string
	LineNumber int
	ClassName  string
	MethodName string
}

// NewCaller payload
func NewCaller(fileName string, lineNumber int, className string, methodName string) Caller {
	return Caller{
		fileName,
		lineNumber,
		className,
		methodName,
	}
}

// GetType of this payload
func (c Caller) GetType() string {
	return "caller"
}

// GetContent from this payload
func (c Caller) GetContent() map[string]interface{} {
	return map[string]interface{}{
		"frame": map[string]interface{}{
			"file_name":   c.FileName,
			"line_number": c.LineNumber,
			"class":       c.ClassName,
			"method":      c.MethodName,
		},
	}
}

// ModifiesLastMessage by this payload?
func (c Caller) ModifiesLastMessage() bool {
	return false
}
