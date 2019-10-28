# pasvorto

Pasvorto is a library written in pure Go providing a set of functions that allow you to generate a random password. 
Supports password rules and length customization. Password strength is detected by [zxcvbn](https://github.com/dropbox/zxcvbn) and is highly reliable.

## Quick Start

Download and install

```bash
go get github.com/marspere/pasvorto
```

```bash
# assume the following codes in example.go file
$ cat example.go
```

```go
package main

import (
	"fmt"
	"github.com/marspere/pasvorto"
)

func main() {
	pwd := pasvorto.Generate(12, pasvorto.Letter, pasvorto.Digital)
	fmt.Println("value:", pwd.Value)
	fmt.Println("score:", pwd.Score)
}
```

```bash
# run example.go
$ go run example.go
```

## Contributing

If you’d like to propose a change please ensure the following:

- All existing tests are passing.
- There are tests in the test suite that cover the changes you’re making.
- You have added documentation strings (in English) to (at least) the public functions you’ve added or modified.