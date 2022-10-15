package results

import (
	"log"
)

// Use this struct to wrap the result of your operation
type Result[T any] struct {
	value T
	err   error
}

// Wrap your operations with is function
// Example:
//
// From:
//     value, err:= do_something(...)
// To:
//     value:= GetResult(do_something(...))
func GetResult[T any](arr T, err error) *Result[T] {
	return &Result[T]{value: arr, err: err}
}

// Use this method to check extract the result. 
func (result *Result[T]) Unwrap() T {
	if !result.IsOk(){
       panic("Tried to access a nil value. You should check that the result is ok before accessing the it:\n" + result.err.Error())
	}
	return result.value
}

// Use this method to check extract the result or return a default value
func (result *Result[T]) Unwrap_Or(val T) T {
	if result.IsOk() {
		return result.value
	} else {
		return val
	}
}

// Use this method to check if the result's value is not nil 
func (result *Result[T]) IsOk() bool {
	return result.err == nil
}

// Use this method to either get the result or exit if an error occurs
func (result *Result[T]) GetOrExit() T {
	if !result.IsOk() {
		log.Fatal(result.err)
	}
	return result.Unwrap()
}

// Use this method to either get the result or panic if an error occurs
func (result *Result[T]) GetOrPanic() T {
	if result.err != nil {
		panic(result.err)
	}
	return result.value
}
