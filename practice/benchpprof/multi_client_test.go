package benchpprof

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestX(t *testing.T) {
	for i := 0; i < 50000; i++ {
		res, err := http.Get("http://localhost:8081/multi")
		if err != nil {
			fmt.Println("error: " + err.Error())
		}
		fmt.Println(res)
	}
	time.Sleep(time.Duration(2000 * time.Second))
}
