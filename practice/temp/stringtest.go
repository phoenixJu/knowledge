package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := int64(345553535)
	println(string(a))
	//string是一种类型，int型的无法强制转换成为string，要通过sprintf相关的格式化输出
	println(fmt.Sprint(a))
	println(strconv.FormatInt(a, 10))
}
