package main

import (
	"html/template"
	"os"
)

func main(){
	tmpl, err := template.New("htmlCode").Parse(`{{define "X"}}Hello, {{.}},{{.}}!{{end}}`)
	if err != nil{
		println(err.Error())
		return
	}
	err = tmpl.ExecuteTemplate(os.Stdout,"X", "laozhu")
}