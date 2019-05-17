package main

/**
回溯法最常见的N皇后问题
回溯法本质上是枚举所有情况，对于每种情况如果遇到阻碍，就回到上一步继续尝试，
这种方式通常使用递归
来实现
最常见的是8皇后问题，8皇后有解92个，如果考虑旋转变化，独立解12个
下表给出了n皇后问题的解的个数包括独立解U（OEIS中的数列A002562）以及互不相同的解D（OEIS中的数列A000170）的个数：
n	1	2	3	4	5	6	7	8	9	10	11	12	13	14	..
U	1	0	0	1	2	1	6	12	46	92	341	1,787	9,233	45,752	..
D	1	0	0	2	10	4	40	92	352	724	2,680	14,200	73,712	365,596	..
 */

import "fmt"

const N = 8

var result = make([]int, N)

func init() {
	for _, i := range result {
		result[i] = -1
	}
}
func CallNQueen(row int) {
	// 如果N行都摆放好了打印结果
	if row == N {
		PrintQueen()
		return
	}
	for column := 0; column < N; column++ {
		// 对于每一行，有N列可以尝试摆放
		if IsCheckOk(row, column) {
			// 如果合适，row行摆在column列处
			result[row] = column
			// 摆放下一行
			CallNQueen(row + 1)
		}
	}
}
func IsCheckOk(row int, column int) bool {
	leftUp := column - 1
	rightUp := column + 1
	for i := row - 1; i >= 0; i-- { // 逐行往上考察每一行
		if result[i] == column { // 第i行的column列有棋子吗
			return false
		}
		if leftUp >= 0 { // bug点：包含0，考察左上对角线，第i行左上角有棋子吗
			if result[i] == leftUp {
				return false
			}
		}
		if rightUp < N { // 考察右上对角线，第i行右上角有棋子吗？
			if result[i] == rightUp {
				return false
			}
		}
		leftUp--
		rightUp++
	}
	return true
}

func PrintQueen() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if result[i] == j {
				fmt.Print("Q")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	// 从第一行开始遍历
	CallNQueen(0)
}
