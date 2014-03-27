# iso8601 #

iso8601 is a simple Go package for encoding `time.Time` in JSON in ISO
8601 format, without subsecond resolution or time zone info.

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/joeshaw/iso8601"
)

func main() {
	t := time.Now()

	// Standard JSON format
	// "2014-03-25T16:15:25.701623113-04:00"
	data, _ := json.Marshal(t)
	fmt.Println(string(data))

	// ISO8601 JSON format
	// "2014-03-25T16:15:25"
	data, _ = json.Marshal(iso8601.Time(t))
	fmt.Println(string(data))

	// Output after decoding back to go.  Note the loss of
	// precision and time zone info.
	// 2014-03-25 16:15:25 +0000 +0000
	var t2 iso8601.Time
	json.Unmarshal(data, &t2)
	fmt.Println(t2)
}
```
