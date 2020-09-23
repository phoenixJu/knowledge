package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type MiddleWare func(handler http.HandlerFunc) http.HandlerFunc

func Logging() MiddleWare {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()
			f(w, r)
			return
		}
	}
}

func Test() MiddleWare {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			defer func() { log.Println(r.URL.Path, "test is ok") }()
			f(w, r)
			return
		}
	}
}

func Method(m string) MiddleWare {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			if request.Method != m {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			}
			handler(writer, request)
		}
	}
}

func Chain(f http.HandlerFunc, middlewares ...MiddleWare) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Test(), Logging()))
	_ = http.ListenAndServe(":8080", nil)
}
