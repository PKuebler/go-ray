# Ray `WIP`

This package can be used to send messages to the [Ray app](https://myray.app/).

## Usage

All functions are applicable to the `ray.RayEntry` object or directly to the package (`ray.Caller()`). In this case the `ray.DefaultClient` is used.

```golang
package main

import (
    ray "github.com/pkuebler/go-ray"
)

func main() {
    // default object
    ray.Ray("my message")

    // own entry
	entry := ray.NewRayEntry(ray.DefaultClient, "")
	entry.TypeName(int64(4343))

    // own client inside a docker container
    client := ray.NewRayClient("host.docker.internal", 23517)
	entry := ray.NewRayEntry(client, "")
    
    // reuse ray entry (settings)
    greenMessages := ray.WithColor("green")
    greenMessages.Ray("my message")

    // change setting on current
    greenMessages.SetColor("red")

    // trigger system notification
    greenMessages.Notify("hi")

    // dump request
	req, err := http.NewRequest("POST", "/health-check", strings.NewReader("{\"fieldA\": 234}"))
	if err != nil {
        panic(err)
	}
	ray.Request(req)

    // dump response
    resp, err := http.Get("http://localhost:8080")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
	ray.Reader(resp.Body)
}
```

## Size / Sort

Messages can be assigned to a color to group them.

```golang
// set on current entry object
greenMessages.SetColor("red")
greenMessages.SetSize("lg")

// get a new entry object
redMessages := greenMessages.WithColor("red")
redMessages := greenMessages.WithSize("lg")
```

## Chainable

All functions return a ray.Entry object and can be concatenated.

```golang
// set on current entry object
greenMessages.SetColor("red").Ray("my notice").TypeName(int64(334))
```

## Docker

In a Docker container, `host.docker.internal` must be defined as the host.

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.