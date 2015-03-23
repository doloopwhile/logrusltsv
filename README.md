# logrus-ltsv
[LTSV](http://ltsv.org/) Formatter for [Logrus(github.com/Sirupsen/logrus)](https://github.com/Sirupsen/logrus)

# Usage
```go
package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/doloopwhile/logrusltsv"
)

func main() {
	log.SetFormatter(&logrusltsv.Formatter{})

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A walrus appears")

	// => time:2015-03-23T12:24:35+09:00\tlevel:info\tmsg:A walrus appears\tanimal:walrus\tsize:10
}
```

