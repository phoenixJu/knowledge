package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println(time.Now().Add(-1 * time.Hour).Unix())
}