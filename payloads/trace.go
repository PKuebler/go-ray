package payloads

type Trace struct {
	FileName   string
	LineNumber int
	ClassName  string
	MethodName string
}

func NewTrace() Trace {
	return Trace{
		"filename.go",
		203,
		"className",
		"methodName",
	}
}

func (t Trace) GetType() string {
	return "trace"
}

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

func (t Trace) ModifiesLastMessage() bool {
	return false
}
