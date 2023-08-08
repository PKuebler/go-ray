package ray

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"net/http/httputil"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/google/uuid"
	"github.com/pkuebler/go-ray/payloads"
)

// RayEntry is the final or intermediate ray entry. It contains color, size and
// the final uuid. These objects can be reused to avoid having to redefine font
// size or color.
type RayEntry struct {
	Client *RayClient

	// Color that can be used are green, orange, red, blue, purple and gray. An
	// empty string removes the color.
	Color string

	// Size that can be used are sm, lg or an empty string.
	Size string

	// UUID of the last message
	UUID string
}

// NewRayEntry creates a new RayEntry with default settings. If the UUID is
// empty, a new one is generated. The UUID is overwritten with each new message.
func NewRayEntry(client *RayClient, id string) *RayEntry {
	if id == "" {
		id = uuid.New().String()
	}

	return &RayEntry{
		Client: client,
		Color:  "",
		Size:   "",
		UUID:   id,
	}
}

// sendRequest creates a Ray Message with Origin, Color, Size and the Payload
// and sends it to Ray via the Client.Send function.
func (r *RayEntry) sendRequest(payload payloads.Payload) *RayEntry {
	loads := []map[string]interface{}{}

	file, no := r.caller(3)

	loads = append(loads, map[string]interface{}{
		"type":    payload.GetType(),
		"content": payload.GetContent(),
		"origin":  payloads.NewOrigin(file, no),
	})

	// color
	if r.Color != "" {
		colorPayload := payloads.NewColor(r.Color)
		loads = append(loads, map[string]interface{}{
			"type":    colorPayload.GetType(),
			"content": colorPayload.GetContent(),
		})
	}

	// size
	if r.Size != "" {
		sizePayload := payloads.NewSize(r.Size)
		loads = append(loads, map[string]interface{}{
			"type":    sizePayload.GetType(),
			"content": sizePayload.GetContent(),
		})
	}

	id := r.UUID
	if !payload.ModifiesLastMessage() {
		id = uuid.New().String()
	}

	msg := Message{
		UUID:     id,
		Payloads: loads,
		Meta: map[string]string{
			"golang_version": fmt.Sprintf(
				"%s, %s, %s",
				runtime.Version(),
				runtime.GOOS,
				runtime.GOARCH,
			),
		},
	}

	r.Client.Send(msg)

	return r
}

// WithColor creates a new RayEntry with the passed color.
func (r *RayEntry) WithColor(color string) *RayEntry {
	return r.Clone().SetColor(color)
}

// WithSize creates a new RayEntry with the passed size.
func (r *RayEntry) WithSize(size string) *RayEntry {
	return r.Clone().SetSize(size)
}

// SetColor sets the color of the current RayEntry.
func (r *RayEntry) SetColor(color string) *RayEntry {
	r.Color = color

	return r
}

// SetSize sets the size of the current RayEntry.
func (r *RayEntry) SetSize(size string) *RayEntry {
	r.Size = size

	return r
}

// NewScreen creates a new screen with optional name. All following entries
// will be displayed on this screen.
func (r *RayEntry) NewScreen(name string) *RayEntry {
	payload := payloads.NewScreen(name)

	return r.sendRequest(payload)
}

// ClearScreen creates a new screen with empty entry list.
func (r *RayEntry) ClearScreen() *RayEntry {
	return r.NewScreen("")
}

// Remove xxx
func (r *RayEntry) Remove() *RayEntry {
	payload := payloads.NewRemove()

	return r.sendRequest(payload)
}

// Hide xxx
func (r *RayEntry) Hide() *RayEntry {
	payload := payloads.NewHide()

	return r.sendRequest(payload)
}

// Trace sent the trace as custom message to Ray. Unfortunately, Ray's trace
// message is not easy to fill with the go trace.
func (r *RayEntry) Trace() *RayEntry {
	s := strings.Replace(string(debug.Stack()), "\t", "&nbsp;&nbsp;&nbsp;&nbsp;", -1)
	s = strings.Replace(s, "\n", "<br />", -1)
	payload := payloads.NewCustom("Trace", s)

	return r.sendRequest(payload)
}

// Caller sends the current position in the code to Ray.
func (r *RayEntry) Caller() *RayEntry {
	file, no := r.caller(2)

	payload := payloads.NewCaller(file, no, "", "")

	return r.sendRequest(payload)
}

// Notify triggers a system notification about Ray.
func (r *RayEntry) Notify(text string) *RayEntry {
	payload := payloads.NewNotify(text)

	return r.sendRequest(payload)
}

// JSON convert and send a interface to json string
func (r *RayEntry) JSON(obj interface{}) *RayEntry {
	payload := payloads.NewJSON(obj)

	return r.sendRequest(payload)
}

// JSONString send a string with json as json string
func (r *RayEntry) JSONString(jsonString string) *RayEntry {
	payload := payloads.NewJSONString(jsonString)

	return r.sendRequest(payload)
}

// Ray send parameters to ray
func (r *RayEntry) Ray(obj ...interface{}) *RayEntry {
	payload := payloads.NewLog(obj)

	return r.sendRequest(payload)
}

// TypeName send the type name to ray as custom message
func (r *RayEntry) TypeName(obj interface{}) *RayEntry {
	typeName := fmt.Sprintf("%T\n", obj)

	payload := payloads.NewCustom("TypeName", typeName)

	return r.sendRequest(payload)
}

// DumpStruct print the struct at ray. Requires a pointer.
func (r *RayEntry) DumpStruct(obj interface{}) *RayEntry {
	v := reflect.ValueOf(obj)
	element := v.Elem()
	typeOfT := element.Type()

	dump := ""

	for i := 0; i < element.NumField(); i++ {
		f := element.Field(i)
		value := r.format(fmt.Sprintf("%v", f.Interface()))
		dump = fmt.Sprintf(
			"%s\n\t%s <span class=sf-dump-key>%s</span>:\t <span class=sf-dump-str>%s</span>",
			dump,
			typeOfT.Field(i).Name,
			f.Type(),
			value,
		)
	}

	payload := payloads.NewLog([]interface{}{
		fmt.Sprintf(
			"<pre class=sf-dump data-indent-pad=\"  \"><span class=sf-dump-note>%s</span> {%s\n}\n</pre>",
			typeOfT,
			dump,
		),
	})

	return r.sendRequest(payload)
}

// Charles print icons at Ray
func (r *RayEntry) Charles() *RayEntry {
	payload := payloads.NewLog([]interface{}{
		"🎶 🎹 🎷 🕺",
	})

	return r.sendRequest(payload)
}

// Custom message with content and an optional label in Ray.
func (r *RayEntry) Custom(content string, label string) *RayEntry {
	payload := payloads.NewCustom(label, content)

	return r.sendRequest(payload)
}

// Clone RayEntry
func (r *RayEntry) Clone() *RayEntry {
	return NewRayEntry(r.Client, "").SetColor(r.Color).SetSize(r.Size)
}

// Request dump a golang request at Ray
func (r *RayEntry) Request(req *http.Request) *RayEntry {
	requestDump, _ := httputil.DumpRequest(req, true)

	s := r.format(string(requestDump))
	payload := payloads.NewCustom("Request", s)

	return r.sendRequest(payload)
}

// Reader dump e.g. request response body
func (r *RayEntry) Reader(reader io.Reader) *RayEntry {
	bodyBytes, err := io.ReadAll(reader)
	if err != nil {
		payload := payloads.NewCustom("Reader", err.Error())

		return r.sendRequest(payload)
	}

	s := r.format(string(bodyBytes))
	s = fmt.Sprintf("<pre class=sf-dump data-indent-pad=\"  \">%s</pre>", s)
	payload := payloads.NewCustom("Reader", s)

	return r.sendRequest(payload)
}

// format escapes the string and converts line breaks to html.
func (r *RayEntry) format(s string) string {
	s = html.EscapeString(s)
	return strings.ReplaceAll(s, "\n", "<br />")
}

// caller returns the file and line
func (r *RayEntry) caller(level int) (string, int) {
	_, file, no, ok := runtime.Caller(level)
	if ok {
		return file, no
	}
	return "", 0
}
