package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("fix")
	if nil != err{
		panic("create file fail")
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprintln(w, "this is a file message")
	w.Flush()
}
