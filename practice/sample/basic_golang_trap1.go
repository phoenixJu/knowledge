package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

func main()  {
	// string 是不可修改的，正确的修改姿势
	s := "Hello你好"
	rune := []rune(s)
	rune[0] = '祝'
	println(string(rune))
	// string的索引操作符返回的是一个byte值，不是字符
	time.Sleep(3 * time.Second)
	x:= "ascii"
	fmt.Println(x[0])
	// go 内建函数len返回字符串的byte数量
	char := "♥"
	fmt.Println(len(char))
	fmt.Println(utf8.RuneCountInString(char))
	char = "é"
	fmt.Println(len(char))    // 3
	fmt.Println(utf8.RuneCountInString(char))    // 2
	fmt.Println("cafe\u0301")    // café    // 法文的 cafe，实际上是两个 rune 的组合
	// string range 迭代会尝试将string翻译为utf8，无效的都会乱码
	data := "A\xfe\x02\xff\x04"
	for _, v := range data {
		fmt.Printf("%#x ", v)    // 0x41 0xfffd 0x2 0xfffd 0x4    // 错误
	}
	for _, v := range []byte(data) {
		fmt.Printf("%#x ", v)    // 0x41 0xfe 0x2 0xff 0x4    // 正确
	}
}
