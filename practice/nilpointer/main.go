package main

import (
	"fmt"
)

func Foo() error {
	var err error
	// …
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)        // <nil>
	fmt.Println(err == nil) // false
}
