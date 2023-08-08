package ray

import (
	"bytes"
	"net/http"
	"strings"
	"testing"
)

type TestJsonStruct struct {
	FieldA string `json:"fieldA"`
	FIeldB string `json:"fieldB"`
}

func Test(t *testing.T) {
	t.Parallel()

	entry := NewRayEntry(DefaultClient, "")
	//	entry.WithColor("green").Notify("huhu")
	entry.NewScreen("my screen")
	entry.TypeName(int64(4343))
	entry.Ray([]interface{}{"huhu", "Jo"})
	entry.WithColor("green").Caller()
	entry.Charles()
	entry.Trace()
	entry.JSONString("{\"fieldA\": 234}")
	entry.JSON(TestJsonStruct{
		FieldA: "hallo",
		FIeldB: "welt",
	})
	entry.DumpStruct(&TestJsonStruct{
		FieldA: "hallo",
		FIeldB: "welt",
	})

	req, err := http.NewRequest("POST", "/health-check", strings.NewReader("{\"fieldA\": 234}"))
	if err != nil {
		t.Fatal(err)
	}
	entry.Request(req)

	buf := bytes.NewBufferString("öjk<html>\n<bla>\n</bla></html>")
	entry.Reader(buf)
}
