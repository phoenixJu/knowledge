package main

import "fmt"

var mapA = make(map[int64]interface{}, 0)
func main()  {
	// reflect中最基本的两个interface是Type和Value，分别通过TypeOf和ValueOf获得
	// 获取的是实际类型或者值
	//a := 3888888888888888888
	//fmt.Println(reflect.TypeOf(a))
    //var x interface{}
	//fmt.Println(reflect.TypeOf(x))
	//fmt.Println(reflect.ValueOf(x))
	//type  Person struct{
	//	Name string
	//	Age int
	//}
	//p := &Person{"sprzhing", 33,}
	//fmt.Println(reflect.TypeOf(p))
	//fmt.Println(reflect.ValueOf(p))
	//p2 := Person{"qwj", 33,}
	//fmt.Println(reflect.TypeOf(p2))
	//fmt.Println(reflect.ValueOf(p2))
	////
	//pi := reflect.ValueOf(p2).Interface()
	//fmt.Println(pi)
	if v, ok := mapA[int64(83)].([]int64);!ok{
		fmt.Print("ok, hello")
		fmt.Printf("%v", v)
	}
}
