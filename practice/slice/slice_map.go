package main

import "fmt"

func main()  {
	a := make(map[string][]int64, 0)
	i := 0
	for i <= 20{
		a["IN"] = append(a["IN"], int64(i))
		i++
	}
	fmt.Println("%v", a)
	fmt.Println(len(a["BR"]))
	fmt.Println(len(a["IN"]))
}
