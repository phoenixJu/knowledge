package main

import (
	"encoding/json"
	"fmt"
)

type Deal struct {
	Name *string
	DealItem []*D
	ID int64
	Sig string
}
type D struct {
	I string
	V int64
}
func main()  {
	name := "laozhu"
	items:= make([]*D, 0)
	dg := &D{
		I: "you",
		V: int64(1),
	}
	items = append(items, dg)
	d := &Deal{
		Name: &name,
		ID : int64(666),
		Sig: "Vigo",
		DealItem: items,
	}
	bytes, _ := json.Marshal(d)
	fmt.Printf("result : %v", string(bytes))
	str := `{"Name":"laozhu","DealItem":[{"I":"you","V":1}],"ID":666,"Sig":"Vigo"}`
	var r Deal
	json.Unmarshal([]byte(str), &r )
	fmt.Printf("un res : %v, %v", *r.Name, *r.DealItem[0])
}
