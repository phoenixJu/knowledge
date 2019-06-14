package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Employee struct {
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	Company string `json:"company"`
}

var EmployeeDB map[string]*Employee

func init() {
	EmployeeDB = make(map[string]*Employee)
	EmployeeDB["Mike"] = &Employee{
		Name:    "Mike",
		Age:     32,
		Company: "baidu",
	}
	EmployeeDB["Sprzhing"] = &Employee{
		Name:    "Sprzhing",
		Age:     33,
		Company: "Toutiao",
	}

}
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Welcome !\n")
}

func Name(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "your name is %s\n", ps.ByName("name"))
}

func Employees(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if emp, ok := EmployeeDB[ps.ByName("name")]; ok {
		b, _ := json.Marshal(emp)
		fmt.Fprintln(w, string(b))
		return
	} else {
		fmt.Fprintln(w,"name not found!")
		return
	}

}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/getname/:name", Name)
	router.GET("/employee/:name", Employees)
	log.Fatal(http.ListenAndServe(":8080", router))
}
