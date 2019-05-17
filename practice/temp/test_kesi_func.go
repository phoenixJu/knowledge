package main

import (
	"code.byted.org/anote/anote_crawl/src/utils"
	"code.byted.org/gopkg/logs"
	"fmt"
)

func main() {
	chat(fmt.Sprintf("Stop %v crawler on %v\n %# v", 3, utils.GetLocalIP(), 6))
}

func chat(message string)  {
	logs.Info("hello: %v", message)
}
