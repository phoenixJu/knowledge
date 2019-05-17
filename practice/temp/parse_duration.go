package main

import (
	"fmt"
	"time"
)

func main() {
	format := "01/02/2006"
	toFormat := "2006-01-02"
	tm, _ := time.Parse(format, "01/12/2017")
	fmt.Print(tm.Format(toFormat))
}
