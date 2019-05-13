package main

import (
	"fmt"
)
/**
0-1 背包其实是将选取的N个物品分为N个阶段，每个阶段决定是或者否选择该物品，通过回溯可以
暴力解决比较小的多阶段最优化问题
 */
var maxW = 0

func Pick(i int, existW int, items []int, n int, w int) {
	// i 第i次；existW 已经选择的总重量；items 物体重量存储位置；n 总共多少个物品；w 背包重量
	if existW == w || i == n {
		if existW > maxW {
			maxW = existW
		}
		return
	}
	// 不选择第i个
	Pick(i+1, existW, items, n, w)
	if existW+items[i] <= w {
		Pick(i+1, existW + items[i], items, n, w)
	}
}

func main() {
	items := []int{1, 3, 4, 13, 5, 13, 51, 41, 33, 89 }
	n := 10
	w := 15
	Pick(0, 0, items, n, w)
	fmt.Println( maxW)
}
