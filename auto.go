package ray

import (
	"io"
	"net/http"
)

var (
	// DefaultClient to ray app
	DefaultClient *RayClient = NewRayClient("localhost", 23517)
	// DefaultRayEntry with current settings
	DefaultRayEntry *RayEntry = NewRayEntry(DefaultClient, "")
)

// SetColor sets the color of the default RayEntry.
func SetColor(color string) *RayEntry {
	return DefaultRayEntry.SetColor(color)
}

// SetSize sets the size of the default RayEntry.
func SetSize(size string) *RayEntry {
	return DefaultRayEntry.SetSize(size)
}

// NewScreen creates a new screen with optional name. All following entries
// will be displayed on this screen.
func NewScreen(screenName string) *RayEntry {
	return DefaultRayEntry.NewScreen(screenName)
}

// ClearScreen creates a new screen with empty entry list.
func ClearScreen() *RayEntry {
	return DefaultRayEntry.NewScreen("")
}

// Trace sent the trace as custom message to Ray. Unfortunately, Ray's trace
// message is not easy to fill with the go trace.
func Trace() *RayEntry {
	return DefaultRayEntry.Trace()
}

// Caller sends the current position in the code to Ray.
func Caller() *RayEntry {
	return DefaultRayEntry.Caller()
}

// Notify triggers a system notification about Ray.
func Notify(text string) *RayEntry {
	return DefaultRayEntry.Notify(text)
}

// JSON convert and send a interface to json string
func JSON(obj interface{}) *RayEntry {
	return DefaultRayEntry.JSON(obj)
}

// JSONString send a string with json as json string
func JSONString(jsonString string) *RayEntry {
	return DefaultRayEntry.JSONString(jsonString)
}

// Ray send parameters to ray
func Ray(objs ...interface{}) *RayEntry {
	return DefaultRayEntry.Ray(objs)
}

// TypeName send the type name to ray as custom message
func TypeName(obj interface{}) *RayEntry {
	return DefaultRayEntry.TypeName(obj)
}

// DumpStruct print the struct at ray. Requires a pointer.
func DumpStruct(obj interface{}) *RayEntry {
	return DefaultRayEntry.DumpStruct(obj)
}

// Charles print icons at Ray
func Charles() *RayEntry {
	return DefaultRayEntry.Charles()
}

// Custom message with content and an optional label in Ray.
func Custom(content string, label string) *RayEntry {
	return DefaultRayEntry.Custom(content, label)
}

// Request dump a golang request at Ray
func Request(req *http.Request) *RayEntry {
	return DefaultRayEntry.Request(req)
}

// Reader dump e.g. request response body
func Reader(reader io.Reader) *RayEntry {
	return DefaultRayEntry.Reader(reader)
}
