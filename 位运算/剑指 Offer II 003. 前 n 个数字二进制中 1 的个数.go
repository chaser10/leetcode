package 位运算

func countBits(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i&(i-1)] + 1
	}

	return dp
}

//func countBits(n int) []int {
//	dp := make([]int, n+1)
//	highBit := 0
//	for i := 1; i <= n; i++ {
//		if i&(i-1) == 0 {
//			highBit = i
//		}
//
//		dp[i] = dp[i-highBit] + 1
//	}
//
//	return dp
//}
//
//func countBits(n int) []int {
//	dp := make([]int, n+1)
//	for i := 1; i <= n; i++ {
//		dp[i] = dp[i>>1] + (i & 1)
//	}
//
//	return dp
//}
