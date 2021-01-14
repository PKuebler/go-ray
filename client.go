package ray

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// RayClient contains the configuration to Ray App
type RayClient struct {
	Enabled bool   `json:"enable"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

// NewRayClient create a client
func NewRayClient(host string, port int) *RayClient {
	return &RayClient{
		Enabled: true,
		Host:    host,
		Port:    port,
	}
}

// Send to ray app
func (r *RayClient) Send(msg Message) {
	if !r.Enabled {
		return
	}

	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("http://%s:%d", r.Host, r.Port)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}
