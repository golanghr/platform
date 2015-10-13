# [Golang.hr] Platform logging package
Base logging package (more like helper) for [Golang.hr Platform]. We are using [Logrus].

More info about [Logrus]:

```
Logrus is a structured logger for Go (golang), completely API compatible with
the standard library logger. Godoc.
Please note the Logrus API is not yet stable (pre 1.0).
Logrus itself is completely stable and has been used in many large deployments.
The core API is unlikely to change much but please version control your Logrus
to make sure you aren't fetching latest master on every build.
```

### Example 1 - Running under debug level

```go
// Package main ...
package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/golanghr/platform/logging"
)

var (
	log logging.Logging
)

func init() {
	log = logging.New(map[string]interface{}{
		"formatter": "text",
		"level":     logrus.DebugLevel,
	})

}

func main() {
	log.WithFields(logrus.Fields{"service": "golang.hr"}).Debug("Example logging")
}
```

### License

```
The MIT License (MIT)

Copyright (c) 2015 Golang Croatia

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```


[Golang.hr]: <https://github.com/golanghr>
[Golang.hr Platform]: <https://github.com/golanghr/platform>
[Logrus]: <https://github.com/Sirupsen/logrus>
