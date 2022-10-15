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
