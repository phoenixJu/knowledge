package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"

	"sync"
)

var hostName string

var RemoteUrl = "http://%v:2280"

func main() {
	flag.StringVar(&hostName, "", "10.115.25.189", "host name")
	RemoteUrl = fmt.Sprintf(RemoteUrl, hostName)

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	r.GET("/v1/agent/self", fwd("/v1/agent/self"))
	r.GET("/v1/lookup/name", fwd("/v1/lookup/name"))
	//r.POST("/v1/lookup/conf", fwd("/v1/lookup/conf"))
	r.Run(":2280")
}

type RespCacheEntry struct {
	Header http.Header
	Body   []byte
}

var respCache = sync.Map{}

func fwd(resPath string) func(context *gin.Context) {

	return func(c *gin.Context) {
		url := RemoteUrl + resPath + "?" + c.Request.URL.RawQuery
		url = strings.Replace(url, ".service.-", "", -1)

		var respCacheEntry *RespCacheEntry
		if entry, ok := respCache.Load(url); !ok {
			resp, err := http.Get(url)
			if err != nil {
				c.JSON(500, gin.H{
					"message": fmt.Sprintf("%v", err),
				})
				return
			} else {
				defer resp.Body.Close()
				bytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					c.JSON(500, gin.H{
						"message": fmt.Sprintf("%v", err),
					})
					return
				}

				if resp.StatusCode != http.StatusOK {
					c.Status(resp.StatusCode)
					for k, v := range resp.Header {
						for _, vHeader := range v {
							c.Header(k, vHeader)
						}
					}
					c.Writer.Write(bytes)
					return
				}

				respCacheEntry = &RespCacheEntry{
					Header: resp.Header,
					Body:   bytes,
				}
				respCache.Store(url, respCacheEntry)
			}
		} else {
			respCacheEntry = entry.(*RespCacheEntry)
		}

		c.Status(http.StatusOK)
		for k, v := range respCacheEntry.Header {
			for _, vHeader := range v {
				c.Header(k, vHeader)
			}
		}
		c.Writer.Write(respCacheEntry.Body)
	}
}
