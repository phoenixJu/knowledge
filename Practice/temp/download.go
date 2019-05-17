package main

import (
	"code.byted.org/gopkg/logs"
	"io"
	"net/http"
	"os"
)

func main()  {
	res, _:= http.Get("http://music.163.com/song/media/outer/url?id=421244551.mp3")
	if res.StatusCode >= 300 {
		logs.Error("expect success status code. actual is %v", res.StatusCode)
		return
	}
	defer res.Body.Close()
	// Check that the server actually sent compressed data
	f, err := os.Create("1.mp3")
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
	//switch res.Header.Get("Content-Encoding") {
	//case "gzip":
	//	reader, _ = gzip.NewReader(res.Body)
	//	defer reader.Close()
	//default:
	//	reader = res.Body
	//}
	//
	//res, err = ioutil.ReadAll(reader)
	return
}
