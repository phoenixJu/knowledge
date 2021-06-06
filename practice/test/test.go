package main

import "encoding/json"

type Value struct {
	IsSelected bool
}

type VideoModelExpireType int64

const (
	VideoModelExpireShort  = VideoModelExpireType(0)
	VideoModelExpireMedium = VideoModelExpireType(1)
	VideoModelExpireLong   = VideoModelExpireType(2)
)

func main()  {
	cont := make([][]interface{}, 0)
	row1 := make([]interface{}, 0)
	row1 = append(row1, "hello world")
	row2 := make([]interface{}, 0)
	row2 = append(row2, "nishishishishis")
	cont = append(cont, row1)
	cont = append(cont, row2)
	str, _ := json.Marshal(cont)
	println(string(str))
}


