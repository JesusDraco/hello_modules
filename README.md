# Hello Modules

This is a repo to test Golang Modules

## Installation

Execute the following command:
```bash
go get -u github.com/JesusDraco/hello_modules
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/JesusDraco/hello_modules"
)

func main() {
	message := hello_modules.Hello2("Jesus")
	fmt.Printf(message)
	fmt.Printf(hello_modules.RandomHello(), "Jesus")
}
```