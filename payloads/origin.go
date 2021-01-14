package payloads

// Origin of file
type Origin struct {
	File       string `json:"file"`
	LineNumber int    `json:"line_number"`
}

// NewOrigin creates a origin
func NewOrigin(file string, lineNumber int) Origin {
	return Origin{
		File:       file,
		LineNumber: lineNumber,
	}
}
