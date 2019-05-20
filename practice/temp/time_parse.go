package main

import (
	"fmt"
	"time"
)

func main() {
	re, err := time.Parse(time.RFC3339,"2019-03-26T14:11:16+01:00" )
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(re)
}
