package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		ctx := r.Context()
		fmt.Fprint(os.Stdout, "processing request \n")
		select {
		case <-time.After(10 * time.Second):
			w.Write([]byte("request processed"))
		case <-ctx.Done():
			fmt.Fprint(os.Stderr, "request cancelled\n")
		}
	}))
}
