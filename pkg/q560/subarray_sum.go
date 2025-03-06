package q560

//考虑以 i 结尾, 和为 k 的连续子数组个数

// 1.枚举
// 需要统计符合条件的下标 j 的个数，其中 0≤j≤i 且 [j..i] 这个子数组的和恰好为 k 。
// 如果我们知道 [j,i] 子数组的和，就能 O(1) 推出 [j−1,i] 的和.
func subarraySum(nums []int, k int) int {
	var res int
	for i := range nums {
		// 枚举 先定末位元素下标i, 再从后往前枚举
		// [j...i] 的和为 sum
		// [j-1...i] 的和为 sum = nums[j-1] + ([j...i] 的和)
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				res++
			}
		}
	}
	return res
}

// 1.我们定义 pre[i] 为 [0..i] 里所有数的和，则 pre[i] 可以由 pre[i−1] 递推而来，即：
// pre[i]=pre[i−1]+nums[i]
// 2. 「[j..i] 这个子数组和为 k 」这个条件我们可以转化为
// pre[i]−pre[j−1]==k
// 简单移项可得符合条件的下标 j 需要满足
// pre[j−1]==pre[i]−k
// 这样把求子数组变为了, 处理从 0 到 j-1 的前缀数组.
// 3.那么问题就转化为了「求出前缀和为 sum-k 的个数」。
func subarraySum2(nums []int, k int) int {
	var res int = 0
	ans := make(map[int]int, len(nums))

	// 0是前缀和的初始值
	// 每个前缀和初始值是1
	ans[0] = 1
	sum := 0
	for i := range nums {
		// 前缀和pre[i]
		sum += nums[i]
		if pre, ok := ans[sum-k]; ok {
			res += pre
		}
		ans[sum]++
	}

	return res
}
