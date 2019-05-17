package main

import (
	"encoding/json"
)

type JsonBody struct {
	Name string
	Age  int64
}

func main() {
	//ar:= strings.Index("http://223.99.245.19/amobile.music.tc.qq.com/", "vkey")
	b := JsonBody{
		"zhuhongqquan",
		23,
	}
	a := &JsonBody{
		"liuwentian",
		33,
	}
	jsona, _ := json.Marshal(a)
	jsonb, _ := json.Marshal(b)
	println(string(jsona))
	println(string(jsonb))
}
func gogogo() (res string) {
	r, res := b()
	if r == "" {
		print("fail")
	} else {
		return
	}
	return
}

func b() (r string, b string) {
	return "3", "4"
}
