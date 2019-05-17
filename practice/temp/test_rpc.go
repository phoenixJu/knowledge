package main

import (
	"code.byted.org/anote/item_provider/thrift_gen/anote/item_provider"
	"code.byted.org/anote/rpc/rpc/rpc_item_provider"
	"code.byted.org/gopkg/logs"
	"code.byted.org/kite/kitc"
	"context"
	"encoding/json"
	"time"
)

func main() {
	options := make([]kitc.Option, 0, 4)
	options = append(options,
		kitc.WithConnTimeout(2*time.Second),
		kitc.WithTimeout(2*time.Second))
	options = append(options, kitc.WithHostPort("127.0.0.1:9910"))
	if err := rpc_item_provider.Init(options...); err != nil{
		panic(err)
	}
	result, err := itemProviderClient.Call("ScanOperationContents", context.Background(), item_provider.ScanOperationContentRequest{
		Cursor: "",
		TimeStart: 0,
		TimeEnd:0,
		Count: 2,
	})
	if err != nil{
		logs.Error("error")
		return
	}
	con, err := json.Marshal(result)
	if err != nil{
		logs.Error("error %v", err)
		return
	}
	print(string(con))
}
