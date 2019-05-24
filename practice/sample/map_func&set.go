package main

import "fmt"

type Set map[interface{}]bool

func (s Set) IsExist(k interface{}) bool {
	if v, ok := s[k]; ok {
		return v
	}
	return false
}
func (s Set) Add(k interface{}) {
	s[k] = true
}
func (s Set) Delete(k interface{}) () {
	delete(s, k)
}
func main() {
	var funMap = make(map[int]func(a int) int, 4)
	funMap[0] = func(a int) int {
		return 1
	}
	funMap[1] = func(a int) int {
		return a
	}
	funMap[2] = func(a int) int {
		return a * a
	}
	funMap[3] = func(a int) int {
		return a * a * a
	}
	for _, v := range funMap {
		fmt.Println(v(5))
	}
	var ideaSet Set = make(map[interface{}]bool,0)
	ideaSet.Add(88)
	ideaSet.Add(55)
	println(ideaSet.IsExist(88))
	ideaSet.Delete(88)
	println(ideaSet.IsExist(88))

}
