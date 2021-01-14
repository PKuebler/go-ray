package payloads

type Caller struct {
	FileName   string
	LineNumber int
	ClassName  string
	MethodName string
}

func NewCaller(fileName string, lineNumber int, className string, methodName string) Caller {
	return Caller{
		fileName,
		lineNumber,
		className,
		methodName,
	}
}

func (c Caller) GetType() string {
	return "caller"
}

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

func (c Caller) ModifiesLastMessage() bool {
	return false
}
