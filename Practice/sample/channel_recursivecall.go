package main
//
//import (
//	"context"
//	"time"
//)
//
//func recursiveCall(ctx context.Context, id []byte, initialNodes [] string) {
//	seen := map[string]string{}     //已经见过的姓名
//	request := make(chan string, 3) //3维度检索
//	//输入初始节点
//	go func() {
//		for _, n := range initialNodes {
//			request <- n
//		}
//	}()
//OUT:
//	for {
//		// 循环直到找到人名
//		if data != nil{
//			return
//		}
//		//在新的请求，超市和上层取消请求中select
//		select {
//		case n:= <-request:
//			go func() {
//				response:= s.sendQuery(ctx, n, MethodfindValue, id)
//				select {
//				case <-ctx.Done():
//				case msg:=<-response:
//					seen[responseToNode(response)] = n //更新已经见过的节点信息
//					//加载新节点
//					for _, rn := range LoadNodeInfoFromByte(msg[PayLoadStart:]){
//						mu.Lock()
//						_, ok := seen[rn.HexId()]
//						mu.Unlock()
//						if ok{
//							continue
//						}
//						AddNode(rn)
//						request <- rn
//					}
//				}
//			}()
//			case <-time.After(500 * time.Millisecond):
//				break OUT
//				case <- ctx.Done()
//				break OUT
//		}
//		return
//	}
//}
