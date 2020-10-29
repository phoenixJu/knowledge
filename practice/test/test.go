package main

import (
	"fmt"
)
type Value struct {
	IsSelected bool
}

func main()  {
	var a interface{}

	x := fmt.Sprintf("%v", a)
	fmt.Printf("%v", x)
}
