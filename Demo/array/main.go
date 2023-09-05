package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}))       // false
	//fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2})) // true
	fmt.Println(minimumDeleteSum("delete", "leet"))
}

func minimumDeleteSum(s1, s2 string) int {
	n, m := len(s1), len(s2)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for j := 1; j <= m; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}
	fmt.Println(dp)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			ch1 := float64(s1[i-1])
			ch2 := float64(s2[j-1])
			if ch1 != ch2 {
				dp[i][j] = int(math.Min(ch1+ch2+float64(dp[i-1][j-1]), math.Min(ch1+float64(dp[i-1][j]), ch2+float64(dp[i][j-1]))))
			} else {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]
}

// 按题的描述
// 用两个 max 变量记录左侧最大值和数组最大值，类似快慢指针的思路
func partitionDisjoint(A []int) int {
	leftMax, curMax, edge := A[0], A[0], 1
	for i, a := range A {
		if leftMax > a { // 比左侧最大值还小，左小数组向右拉长
			leftMax = curMax // 目前的最大值同步到 leftMax
			edge = i + 1
		}
		if a > curMax {
			curMax = a
		}
	}
	return edge
}

//卡牌分组 简单
func hasGroupsSizeX(deck []int) bool {
	if len(deck) <= 1 {
		return false
	}
	freqs := make(map[int]int)
	for _, num := range deck {
		freqs[num]++
	}
	counts := make([]int, 0, len(freqs))
	for _, freq := range freqs {
		counts = append(counts, freq)
	}
	for i := 0; i < len(counts)-1; i++ {
		if gcd(counts[i], counts[i+1]) == 1 { //不能分组
			return false
		}
	}
	return true
}

//辗转相除法求最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
