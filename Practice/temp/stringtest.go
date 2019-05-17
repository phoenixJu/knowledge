package main

import (
	"code.byted.org/anote/lib/util"
	"strconv"
)

func main() {
	a := int64(345553535)
	println(string(a))
	//string是一种类型，int型的无法强制转换成为string，要通过sprintf相关的格式化输出
	println(an_util.ToString(a))
	println(an_util.EncodeInt64Cursor(a))
	println(strconv.FormatInt(a, 10))
}
