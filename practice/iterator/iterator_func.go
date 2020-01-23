package main

import "fmt"

type Ints []int
// Ints和index目前类似在Lua中，以Upvalue被闭包捕获
func (i Ints) Iterator() func() (int, bool) {
	index := 0
	return func() (val int, ok bool) {
		if index >= len(i) {
			return
		}
		val, ok = i[index], true
		index++
		return
	}
}

func main1() {
	ints := Ints{1, 2, 5, 8, 13}
	it := ints.Iterator()
	for {
		val, ok := it()
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
