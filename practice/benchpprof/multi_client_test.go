package benchpprof

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"testing"
	"time"
)

func TestX(t *testing.T) {
	f, err := os.Create("./multi_cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil { //监控cpu
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()
	for i := 0; i< 50000; i++{
		res, err := http.Get("http://localhost:8081/multi")
		if err != nil{
			fmt.Println("error: " + err.Error())
		}
		fmt.Println(res)
	}
	time.Sleep(10000* time.Duration(20 * time.Second))
}
