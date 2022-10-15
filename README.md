# Go-Results

## Use this package for error handling using Go
```
Error handling has always been an issue in Go.
The Goal of this package is to make error handling in Go easier and limit access to unexpected nil values. 
Similar to rust's idea of using a result struct, you wrap the result of your operation
with GetResult to create a Result struct then use the available methods to continue.

This package is still under construction.

```
### Usage Example

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	file := GetResult(os.ReadFile("result.go")).GetOrPanic()

	fmt.Println(string(file))

	file2 := GetResult(os.ReadFile("results.go")).Unwrap_Or(nil)
	fmt.Println("file is nil =>", file2 == nil)

	if GetResult(os.ReadFile("resul.go")).IsOk() {
		fmt.Println("file exists")
	} else {
		fmt.Println("file does not exists")
	}

}

```